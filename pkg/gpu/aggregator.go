// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

//go:build linux_bpf

package gpu

import (
	"github.com/DataDog/datadog-agent/pkg/collector/corechecks/gpu/model"
)

const nsecPerSec = 1e9

// aggregator is responsible for receiving stream-level data for one process and aggregating it into metrics
type aggregator struct {
	// totalThreadSecondsUsed is the total amount of thread-seconds used by the
	// GPU in the current interval
	totalThreadSecondsUsed float64

	// lastCheckKtime is the last kernel time the processor was checked,
	// required to take into account only data from the current interval.
	lastCheckKtime int64

	// measuredIntervalNs is the interval between the last two checks, in
	// nanoseconds, required to compute thread seconds used/available for
	// utilization percentages
	measuredIntervalNs int64

	// currentAllocs is the list of current (active) memory allocations
	currentAllocs []*memoryAllocation

	// pastAllocs is the list of past (freed) memory allocations
	pastAllocs []*memoryAllocation

	// utilizationNormFactor is the factor to normalize the utilization by, to
	// account for the fact that we might have more kernels enqueued than the
	// GPU can run in parallel. This factor allows distributing the utilization
	// over all the streams that were active during the interval.
	utilizationNormFactor float64

	// processEnded is true if the process has ended and this aggregator should be deleted
	processEnded bool

	// sysCtx is the system context with global GPU-system data
	sysCtx *systemContext
}

func newAggregator(sysCtx *systemContext) *aggregator {
	return &aggregator{
		sysCtx: sysCtx,
	}
}

// processKernelSpan takes a kernel span and computes the thread-seconds used by it during
// the interval
func (agg *aggregator) processKernelSpan(span *kernelSpan) {
	tsStart := int64(span.startKtime)
	tsEnd := int64(span.endKtime)

	// we only want to consider data that was not already processed in the previous interval
	if agg.lastCheckKtime > tsStart {
		tsStart = agg.lastCheckKtime
	}

	durationSec := float64(tsEnd-tsStart) / nsecPerSec
	maxThreads := uint64(agg.sysCtx.maxGpuThreadsPerDevice[0])                                // TODO: MultiGPU support not enabled yet
	agg.totalThreadSecondsUsed += durationSec * float64(min(span.avgThreadCount, maxThreads)) // we can't use more threads than the GPU has
}

// processPastData takes spans/allocations that have already been closed
func (agg *aggregator) processPastData(data *streamData) {
	for _, span := range data.spans {
		agg.processKernelSpan(span)
	}

	agg.pastAllocs = append(agg.pastAllocs, data.allocations...)
}

// processCurrentData takes spans/allocations that are active (e.g., unfreed allocations, running kernels)
func (agg *aggregator) processCurrentData(data *streamData) {
	for _, span := range data.spans {
		agg.processKernelSpan(span)
	}

	agg.currentAllocs = append(agg.currentAllocs, data.allocations...)
}

// getGpuUtilization computes the utilization of the GPU as the average of the
func (agg *aggregator) getGPUUtilization() float64 {
	intervalSecs := float64(agg.measuredIntervalNs) / nsecPerSec
	if intervalSecs > 0 {
		// TODO: MultiGPU support not enabled yet
		availableThreadSeconds := float64(agg.sysCtx.maxGpuThreadsPerDevice[0]) * intervalSecs
		return agg.totalThreadSecondsUsed / availableThreadSeconds
	}

	return 0
}

// setGPUUtilizationNormalizationFactor sets the factor to normalize GPU utilization
func (agg *aggregator) setGPUUtilizationNormalizationFactor(factor float64) {
	agg.utilizationNormFactor = factor
}

// getStats returns the aggregated stats for the process
func (agg *aggregator) getStats() model.PIDStats {
	var stats model.PIDStats

	if agg.measuredIntervalNs > 0 {
		stats.UtilizationPercentage = agg.getGPUUtilization() / agg.utilizationNormFactor
	}

	var memTsBuilder tseriesBuilder

	for _, alloc := range agg.currentAllocs {
		memTsBuilder.AddEventStart(alloc.startKtime, int64(alloc.size))
	}

	for _, alloc := range agg.pastAllocs {
		memTsBuilder.AddEvent(alloc.startKtime, alloc.endKtime, int64(alloc.size))
	}

	lastValue, maxValue := memTsBuilder.GetLastAndMax()
	stats.CurrentMemoryBytes = uint64(lastValue)
	stats.MaxMemoryBytes = uint64(maxValue)

	// Flush the data that we used
	agg.flushProcessedStats()

	return stats
}

func (agg *aggregator) flushProcessedStats() {
	agg.currentAllocs = agg.currentAllocs[:0]
	agg.pastAllocs = agg.pastAllocs[:0]
	agg.totalThreadSecondsUsed = 0
}
