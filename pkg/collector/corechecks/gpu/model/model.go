// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

// Package model contains the model for the GPU check, with types shared between the system-probe GPU probe and
// the gpu core agent check
package model

// PIDStats contains the GPU stats for a given PID
type PIDStats struct {
	UtilizationPercentage float64
	CurrentMemoryBytes    uint64
	MaxMemoryBytes        uint64
}

// GPUStats contains the past and current data for all streams, including kernel spans and allocations.
// This is the data structure that is sent to the agent
type GPUStats struct {
	PIDStats map[uint32]PIDStats
}
