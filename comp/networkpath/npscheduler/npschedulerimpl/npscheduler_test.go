// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

package npschedulerimpl

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/DataDog/datadog-agent/comp/aggregator/demultiplexer/demultiplexerimpl"
	"github.com/DataDog/datadog-agent/comp/core"
	"github.com/DataDog/datadog-agent/comp/core/sysprobeconfig/sysprobeconfigimpl"
	"github.com/DataDog/datadog-agent/comp/forwarder/defaultforwarder"
	"github.com/DataDog/datadog-agent/comp/forwarder/eventplatform/eventplatformimpl"
	"github.com/DataDog/datadog-agent/comp/ndmtmp/forwarder/forwarderimpl"
	"github.com/DataDog/datadog-agent/comp/networkpath/npscheduler"
	utillog "github.com/DataDog/datadog-agent/pkg/util/log"
	"github.com/cihub/seelog"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

// testOptions is a fx collection of common dependencies for all tests
var testOptions = fx.Options(
	Module(),
	forwarderimpl.MockModule(),
	demultiplexerimpl.MockModule(),
	defaultforwarder.MockModule(),
	core.MockBundle(),
	eventplatformimpl.MockModule(),
)

func TestStartServerAndStopNpScheduler(t *testing.T) {
	var component npscheduler.Component
	app := fxtest.New(t, fx.Options(
		testOptions,
		fx.Supply(fx.Annotate(t, fx.As(new(testing.TB)))),
		fx.Replace(sysprobeconfigimpl.MockParams{Overrides: map[string]any{
			"network_path.enabled": true,
		}}),
		fx.Populate(&component),
	))
	npScheduler := component.(*npSchedulerImpl)

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	l, err := seelog.LoggerFromWriterWithMinLevelAndFormat(w, seelog.DebugLvl, "[%LEVEL] %FuncShort: %Msg")
	assert.Nil(t, err)
	utillog.SetupLogger(l, "debug")

	assert.NotNil(t, npScheduler)
	assert.NotNil(t, app)
	assert.False(t, npScheduler.running)
	app.RequireStart()
	assert.True(t, npScheduler.running)
	app.RequireStop()
	assert.False(t, npScheduler.running)

	w.Flush()
	logs := b.String()

	assert.Equal(t, 1, strings.Count(logs, "Start NpScheduler"), logs)
	assert.Equal(t, 1, strings.Count(logs, "Starting listening for pathtests"), logs)
	assert.Equal(t, 1, strings.Count(logs, "Starting flush loop"), logs)
	assert.Equal(t, 1, strings.Count(logs, "Starting workers"), logs)
	assert.Equal(t, 1, strings.Count(logs, "Starting worker #0"), logs)

	assert.Equal(t, 1, strings.Count(logs, "Stopped listening for pathtests"), logs)
	assert.Equal(t, 1, strings.Count(logs, "Stopped flush loop"), logs)
	assert.Equal(t, 1, strings.Count(logs, "Stop NpScheduler"), logs)
}

func Test_newNpSchedulerImpl_defaultConfigs(t *testing.T) {
	var component npscheduler.Component
	app := fxtest.New(t, fx.Options(
		testOptions,
		fx.Supply(fx.Annotate(t, fx.As(new(testing.TB)))),
		fx.Replace(sysprobeconfigimpl.MockParams{Overrides: map[string]any{
			"network_path.enabled": true,
		}}),
		fx.Populate(&component),
	))
	npScheduler := component.(*npSchedulerImpl)

	assert.NotNil(t, npScheduler)
	assert.NotNil(t, app)

	assert.Equal(t, true, npScheduler.enabled)
	assert.Equal(t, 4, npScheduler.workers)
	assert.Equal(t, 1000, cap(npScheduler.pathtestInputChan))
	assert.Equal(t, 1000, cap(npScheduler.pathtestProcessChan))
}

func Test_newNpSchedulerImpl_overrideConfigs(t *testing.T) {
	var component npscheduler.Component
	app := fxtest.New(t, fx.Options(
		testOptions,
		fx.Supply(fx.Annotate(t, fx.As(new(testing.TB)))),
		fx.Replace(sysprobeconfigimpl.MockParams{Overrides: map[string]any{
			"network_path.enabled":           true,
			"network_path.workers":           2,
			"network_path.input_chan_size":   300,
			"network_path.process_chan_size": 400,
		}}),
		fx.Populate(&component),
	))
	npScheduler := component.(*npSchedulerImpl)

	assert.NotNil(t, npScheduler)
	assert.NotNil(t, app)

	assert.Equal(t, true, npScheduler.enabled)
	assert.Equal(t, 2, npScheduler.workers)
	assert.Equal(t, 300, cap(npScheduler.pathtestInputChan))
	assert.Equal(t, 400, cap(npScheduler.pathtestProcessChan))
}
