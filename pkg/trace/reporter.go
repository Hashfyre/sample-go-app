package trace

import (
	"time"

	"github.com/uber/jaeger-client-go/config"
)

// LocalAgentReporter and other reporter types
const (
	LocalAgentReporter = "localagent"
	CollectorReporter  = "collector"
)

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

func NewReporterConfig(reporter Reporter) config.ReporterConfig {
	switch reporter.Type {
	case LocalAgentReporter:
		return newLocalAgentReporterConfig(reporter)
	case CollectorReporter:
		return newCollectorReporterConfig(reporter)
	}
	return config.ReporterConfig{}
}

func newLocalAgentReporterConfig(reporter Reporter) config.ReporterConfig {
	return config.ReporterConfig{
		QueueSize:           reporter.QueueSize,
		BufferFlushInterval: reporter.BufferFlushInterval,
		LogSpans:            reporter.LogSpans,
		LocalAgentHostPort:  reporter.LocalAgent.Host + ":" + reporter.LocalAgent.Port,
	}
}

func newCollectorReporterConfig(reporter Reporter) config.ReporterConfig {
	return config.ReporterConfig{
		QueueSize:           reporter.QueueSize,
		BufferFlushInterval: reporter.BufferFlushInterval,
		LogSpans:            reporter.LogSpans,
		CollectorEndpoint:   reporter.Collector.Endpoint,
		User:                reporter.Collector.User,
		Password:            reporter.Collector.Password,
		HTTPHeaders:         reporter.Collector.Headers,
	}
}
