// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//nolint:revive
package aggregator

import (
	"context"
	"fmt"
	"sync"

	hostMetadataUtils "github.com/DataDog/datadog-agent/comp/metadata/host/hostimpl/hosttags"
	pkgconfigsetup "github.com/DataDog/datadog-agent/pkg/config/setup"

	"github.com/benbjohnson/clock"
)

type HostTagProvider struct {
	hostTags []string
	sync.RWMutex
}

func NewHostTagProvider() *HostTagProvider {
	return newHostTagProviderWithClock(clock.New())
}

func newHostTagProviderWithClock(clock clock.Clock) *HostTagProvider {
	p := &HostTagProvider{
		hostTags: []string{},
	}
	duration := pkgconfigsetup.Datadog().GetDuration("expected_tags_duration")
	fmt.Println("hehexd", duration)
	p.hostTags = append(p.hostTags, hostMetadataUtils.Get(context.TODO(), false, pkgconfigsetup.Datadog()).System...)
	fmt.Println("hehexd2", p.hostTags)
	if pkgconfigsetup.Datadog().GetDuration("expected_tags_duration") > 0 {
		expectedTagsDeadline := pkgconfigsetup.StartTime.Add(duration)
		clock.AfterFunc(expectedTagsDeadline.Sub(clock.Now()), func() {
			p.Lock()
			defer p.Unlock()
			p.hostTags = nil
		})
	}

	return p
}

func (p *HostTagProvider) GetHostTags() []string {
	p.RLock()
	defer p.RUnlock()

	if p.hostTags != nil && len(p.hostTags) > 0 {
		return p.hostTags
	}
	return nil
}
