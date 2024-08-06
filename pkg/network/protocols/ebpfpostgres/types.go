// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

//go:build ignore

package ebpfpostgres

/*
#include "../../ebpf/c/protocols/postgres/types.h"
#include "../../ebpf/c/protocols/classification/defs.h"
*/
import "C"

// This package was created to avoid cyclic imports.

type ConnTuple = C.conn_tuple_t

type EbpfEvent C.postgres_event_t
type EbpfTx C.postgres_transaction_t

const (
	BufferSize = C.POSTGRES_BUFFER_SIZE
)
