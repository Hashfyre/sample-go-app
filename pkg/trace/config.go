package trace

import (
	"github.com/uber/jaeger-client-go/config"
)

type Config struct {
	App      string
	Sampler  Sampler
	Reporter Reporter
}

func (cfg Config) New() config.Configuration {
	samplerConfig := NewSamplerConfig(cfg.Sampler)
	reporterConfig := NewReporterConfig(cfg.Reporter)
	return config.Configuration{
		ServiceName: cfg.App,
		Sampler:     &samplerConfig,
		Reporter:    &reporterConfig,
		// Headers:             {}, // TODO
		// BaggageRestrictions: {}, // TODO
		// Throttler:           {}, //TODO
	}
}
