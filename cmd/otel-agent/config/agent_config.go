// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Package config provides a way to convert the OpenTelemetry Collector configuration to the Datadog Agent configuration.
package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/DataDog/datadog-agent/comp/core/config"
	"github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/exporter/datadogexporter"
	pkgconfigmodel "github.com/DataDog/datadog-agent/pkg/config/model"
	pkgconfigsetup "github.com/DataDog/datadog-agent/pkg/config/setup"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/converter/expandconverter"
	"go.opentelemetry.io/collector/confmap/provider/envprovider"
	"go.opentelemetry.io/collector/confmap/provider/fileprovider"
	"go.opentelemetry.io/collector/confmap/provider/httpprovider"
	"go.opentelemetry.io/collector/confmap/provider/httpsprovider"
	"go.opentelemetry.io/collector/confmap/provider/yamlprovider"
	"go.opentelemetry.io/collector/service"
)

// NewConfigComponent creates a new config component from the given URIs
func NewConfigComponent(ctx context.Context, uris []string) (config.Component, error) {
	// Load the configuration from the fileName
	rs := confmap.ResolverSettings{
		URIs: uris,
		ProviderFactories: []confmap.ProviderFactory{
			fileprovider.NewFactory(),
			envprovider.NewFactory(),
			yamlprovider.NewFactory(),
			httpprovider.NewFactory(),
			httpsprovider.NewFactory(),
		},
		ConverterFactories: []confmap.ConverterFactory{expandconverter.NewFactory()},
	}

	resolver, err := confmap.NewResolver(rs)
	if err != nil {
		return nil, err
	}
	cfg, err := resolver.Resolve(ctx)
	if err != nil {
		return nil, err
	}
	ddc, err := getDDExporterConfig(cfg)
	if err != nil {
		return nil, err
	}
	sc, err := getServiceConfig(cfg)
	if err != nil {
		return nil, err
	}
	site := ddc.API.Site
	apiKey := string(ddc.API.Key)
	// Create a new config
	pkgconfig := pkgconfigmodel.NewConfig("OTel", "DD", strings.NewReplacer(".", "_"))
	// Set Default values
	pkgconfigsetup.InitConfig(pkgconfig)
	pkgconfig.Set("api_key", apiKey, pkgconfigmodel.SourceFile)
	pkgconfig.Set("site", site, pkgconfigmodel.SourceFile)

	pkgconfig.Set("logs_enabled", true, pkgconfigmodel.SourceFile)
	pkgconfig.Set("logs_config.use_compression", true, pkgconfigmodel.SourceFile)
	pkgconfig.Set("log_level", sc.Telemetry.Logs.Level, pkgconfigmodel.SourceFile)
	pkgconfig.Set("apm_config.enabled", true, pkgconfigmodel.SourceFile)
	pkgconfig.Set("apm_config.apm_non_local_traffic", true, pkgconfigmodel.SourceFile)
	return pkgconfig, nil
}

func getServiceConfig(cfg *confmap.Conf) (*service.Config, error) {
	var pipelineConfig *service.Config
	s := cfg.Get("service")
	if s == nil {
		return nil, fmt.Errorf("service config not found")
	}
	smap := s.(map[string]any)
	err := confmap.NewFromStringMap(smap).Unmarshal(&pipelineConfig)
	if err != nil {
		return nil, err
	}
	return pipelineConfig, nil
}

func getDDExporterConfig(cfg *confmap.Conf) (*datadogexporter.Config, error) {
	var configs []*datadogexporter.Config
	var err error
	for k, v := range cfg.ToStringMap() {
		if k != "exporters" {
			continue
		}
		exporters := v.(map[string]any)
		for k, v := range exporters {
			if strings.HasPrefix(k, "datadog") {
				var datadogConfig *datadogexporter.Config
				m := v.(map[string]any)
				err = confmap.NewFromStringMap(m).Unmarshal(&datadogConfig)
				if err != nil {
					return nil, err
				}
				configs = append(configs, datadogConfig)
			}
		}
	}
	if len(configs) == 0 {
		return nil, fmt.Errorf("no datadog exporter found")
	}
	// Check if we have multiple datadog exporters
	// We only support one exporter for now
	// TODO: support multiple exporters
	if len(configs) > 1 {
		return nil, fmt.Errorf("multiple datadog exporters found")
	}

	datadogConfig := configs[0]
	return datadogConfig, nil
}
