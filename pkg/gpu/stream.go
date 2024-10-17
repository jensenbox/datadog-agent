// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

//go:build linux_bpf

package gpu

import (
	"math"

	ddebpf "github.com/DataDog/datadog-agent/pkg/ebpf"
	gpuebpf "github.com/DataDog/datadog-agent/pkg/gpu/ebpf"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// streamHandler is responsible for receiving events from a single CUDA stream and generating
// stats from them.
type streamHandler struct {
	kernelLaunches []gpuebpf.CudaKernelLaunch
	memAllocEvents map[uint64]gpuebpf.CudaMemEvent
	kernelSpans    []*kernelSpan
	allocations    []*memoryAllocation
	processEnded   bool // A marker to indicate that the process has ended, and this handler should be flushed
}

// streamKey is a unique identifier for a CUDA stream
type streamKey struct {
	pid    uint32
	stream uint64
}

// kernelSpan represents a span of time during which one or more kernels were running on a GPU until
// a synchronization event happened
type kernelSpan struct {
	startKtime     uint64 // startKtime is the kernel-time of the moment the first kernel was launched
	endKtime       uint64 // endKtime is the kernel-time of the moment the synchronization event happened
	avgThreadCount uint64 // avgThreadCount is the average number of threads running on the GPU during the span
	numKernels     uint64 // numKernels is the number of kernels that were launched during the span
}

// memoryAllocation represents a memory allocation event
type memoryAllocation struct {
	startKtime uint64 // StartKtime is the kernel-time timestamp of the allocation event
	endKtime   uint64 // EndKtime is the kernel-time timestamp of the deallocation event. If 0, this means the allocation was not deallocated yet
	size       uint64 // Size is the size of the allocation in bytes
	isLeaked   bool   // IsLeaked is true if the allocation was not deallocated
}

// streamData contains kernel spans and allocations for a stream
type streamData struct {
	key         streamKey
	spans       []*kernelSpan
	allocations []*memoryAllocation
}

func newStreamHandler() *streamHandler {
	return &streamHandler{
		memAllocEvents: make(map[uint64]gpuebpf.CudaMemEvent),
	}
}

func (sh *streamHandler) handleKernelLaunch(event *gpuebpf.CudaKernelLaunch) {
	// Copy events, as the memory can be overwritten in the ring buffer after the function returns
	sh.kernelLaunches = append(sh.kernelLaunches, *event)
}

func (sh *streamHandler) handleMemEvent(event *gpuebpf.CudaMemEvent) {
	if event.Type == gpuebpf.CudaMemAlloc {
		sh.memAllocEvents[event.Addr] = *event
		return
	}

	// We only support alloc and free events for now, so if it's not alloc it's free.
	alloc, ok := sh.memAllocEvents[event.Addr]
	if !ok {
		log.Warnf("Invalid free event: %v", event)
		return
	}

	data := memoryAllocation{
		startKtime: alloc.Header.Ktime_ns,
		endKtime:   event.Header.Ktime_ns,
		size:       alloc.Size,
		isLeaked:   false, // Came from a free event, so it's not a leak
	}

	sh.allocations = append(sh.allocations, &data)
	delete(sh.memAllocEvents, event.Addr)
}

func (sh *streamHandler) markSynchronization(ts uint64) {
	span := sh.getCurrentKernelSpan(ts)
	if span == nil {
		return
	}

	sh.kernelSpans = append(sh.kernelSpans, span)

	remainingLaunches := []gpuebpf.CudaKernelLaunch{}
	for _, launch := range sh.kernelLaunches {
		if launch.Header.Ktime_ns >= ts {
			remainingLaunches = append(remainingLaunches, launch)
		}
	}
	sh.kernelLaunches = remainingLaunches
}

func (sh *streamHandler) handleSync(event *gpuebpf.CudaSync) {
	// TODO: Worry about concurrent calls to this?
	sh.markSynchronization(event.Header.Ktime_ns)
}

func (sh *streamHandler) getCurrentKernelSpan(maxTime uint64) *kernelSpan {
	span := kernelSpan{
		startKtime: math.MaxUint64,
		endKtime:   maxTime,
		numKernels: 0,
	}

	for _, launch := range sh.kernelLaunches {
		// Skip launches that happened after the max time we are interested in
		// For example, do not include launches that happened after the synchronization event
		if launch.Header.Ktime_ns >= maxTime {
			continue
		}

		span.startKtime = min(launch.Header.Ktime_ns, span.startKtime)
		span.endKtime = max(launch.Header.Ktime_ns, span.endKtime)
		blockSize := launch.Block_size.X * launch.Block_size.Y * launch.Block_size.Z
		blockCount := launch.Grid_size.X * launch.Grid_size.Y * launch.Grid_size.Z
		span.avgThreadCount += uint64(blockSize) * uint64(blockCount)
		span.numKernels++
	}

	if span.numKernels == 0 {
		return nil
	}

	span.avgThreadCount /= uint64(span.numKernels)

	return &span
}

// getPastData returns all the events that have finished (kernel spans with synchronizations/allocations that have been freed)
// If flush is true, the data will be cleared from the handler
func (sh *streamHandler) getPastData(flush bool) *streamData {
	if len(sh.kernelSpans) == 0 && len(sh.allocations) == 0 {
		return nil
	}

	data := &streamData{
		spans:       sh.kernelSpans,
		allocations: sh.allocations,
	}

	if flush {
		sh.kernelSpans = nil
		sh.allocations = nil
	}

	return data
}

func (sh *streamHandler) getCurrentData(now uint64) *streamData {
	if len(sh.kernelLaunches) == 0 && len(sh.memAllocEvents) == 0 {
		return nil
	}

	data := &streamData{}
	span := sh.getCurrentKernelSpan(now)
	if span != nil {
		data.spans = append(data.spans, span)
	}

	for _, alloc := range sh.memAllocEvents {
		data.allocations = append(data.allocations, &memoryAllocation{
			startKtime: alloc.Header.Ktime_ns,
			endKtime:   0,
			size:       alloc.Size,
			isLeaked:   false,
		})
	}

	return data
}

// markEnd is called when this stream is closed (process exited or stream destroyed).
// A synchronization event will be triggered and all pending events (allocations) will be resolved.
func (sh *streamHandler) markEnd() error {
	nowTs, err := ddebpf.NowNanoseconds()
	if err != nil {
		return err
	}

	sh.processEnded = true
	sh.markSynchronization(uint64(nowTs))

	// Close all allocations. Treat them as leaks, as they weren't freed properly
	for _, alloc := range sh.memAllocEvents {
		data := memoryAllocation{
			startKtime: alloc.Header.Ktime_ns,
			endKtime:   uint64(nowTs),
			size:       alloc.Size,
			isLeaked:   true,
		}
		sh.allocations = append(sh.allocations, &data)
	}

	return nil
}
