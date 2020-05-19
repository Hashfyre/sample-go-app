package trace

import (
	"time"

	"github.com/uber/jaeger-client-go/config"
)

type Constant struct {
	Enabled bool
}

type Probablilistic struct {
	Probability float64
}

type RateLimiting struct {
	SpansPerSec uint
}

type Remote struct {
	Probability     float64
	ServerURL       string
	RefreshInterval time.Duration
	MaxOp           int
	LateBinding     bool
}

type Sampler struct {
	Type          string
	Constant      Constant
	Probabilistic Probablilistic
	RateLimiting  RateLimiting
	Remote        Remote
}

type LocalAgent struct {
	Host string
	Port string
}

type Collector struct {
	Endpoint string
	User     string
	Password string
	Headers  map[string]string
}

type Reporter struct {
	Type                string
	QueueSize           int
	BufferFlushInterval time.Duration
	LogSpans            bool
	LocalAgent          LocalAgent
	Collector           Collector
}

type Config struct {
	App      string
	Sampler  Sampler
	Reporter Reporter
}

func (cfg Config) New() config.Configuration {
	samplerConfig := NewSamplerConfig(cfg.Sampler)
	reporterConfig := NewReporterConfig(cfg.Reporter)
	return config.Configuration{
		Sampler:  &samplerConfig,
		Reporter: &reporterConfig,
		// Headers:             {}, // TODO
		// BaggageRestrictions: {}, // TODO
		// Throttler:           {}, //TODO
	}
}
