package config

import (
	"strconv"
	"time"
)

// JaegarTracer exposes config for jaegar tracer
type JaegarTracer struct {
	AppName                     string
	SamplingProbability         float64 // for Probabilistic sampler
	SamplingRate                uint    // for RateLimiting Sampler, units of Spans per sec
	ReporterQueueSize           int
	ReporterBufferFlushInterval time.Duration
	ReporterLogSpans            bool
	ReporterLocalAgentHost      string
	ReporterLocalAgentPort      string
	ReporterCollectorEndpoint   string
	ReporterCollectorUser       string
	ReporterCollectorPassword   string
}

func initJaegerTracer() (*JaegarTracer, error) {
	appName, err := configCheck("TRACER_APP_NAME")
	if err != nil {
		return nil, err
	}

	samplingProbability, err := configCheck("TRACER_SAMPLING_PROBABILITY")
	if err != nil {
		return nil, err
	}

	sampleProb, err := strconv.ParseFloat(samplingProbability, 64)
	if err != nil {
		return nil, err
	}

	samplingRate, err := configCheck("TRACER_SAMPLING_RATE")
	if err != nil {
		return nil, err
	}

	sampleRate, err := strconv.ParseUint(samplingRate, 10, 32)
	if err != nil {
		return nil, err
	}

	reporterQueueSize, err := configCheck("TRACER_REPORTER_QUEUE_SIZE")
	if err != nil {
		return nil, err
	}

	reporterQSize, err := strconv.Atoi(reporterQueueSize)
	if err != nil {
		return nil, err
	}

	reporterBufferFlushInterval, err := configCheck("TRACER_REPORTER_BUFFER_FLUSH_INTERVAL")
	if err != nil {
		return nil, err
	}

	reporterBufferFlushInt, err := time.ParseDuration(reporterBufferFlushInterval)
	if err != nil {
		return nil, err
	}

	reporterLogSpans, err := configCheck("TRACER_REPORTER_LOG_SPANS")
	if err != nil {
		return nil, err
	}

	logSpans, err := strconv.ParseBool(reporterLogSpans)
	if err != nil {
		return nil, err
	}

	reporterLocalAgentHost, err := configCheck("TRACER_REPORTER_LOCAL_AGENT_HOST")
	if err != nil {
		return nil, err
	}

	reporterLocalAgentPort, err := configCheck("TRACER_REPORTER_LOCAL_AGENT_PORT")
	if err != nil {
		return nil, err
	}

	reporterCollectorEndpoint, err := configCheck("TRACER_REPORTER_COLLECTOR_ENDPOINT")
	if err != nil {
		return nil, err
	}
	reporterCollectorUser, err := configCheck("TRACER_REPORTER_COLLECTOR_USER")
	if err != nil {
		return nil, err
	}
	reporterCollectorPassword, err := configCheck("TRACER_REPORTER_COLLECTOR_PASSWORD")
	if err != nil {
		return nil, err
	}

	return &JaegarTracer{
		AppName:                     appName,
		SamplingProbability:         sampleProb,
		SamplingRate:                uint(sampleRate),
		ReporterQueueSize:           reporterQSize,
		ReporterBufferFlushInterval: reporterBufferFlushInt,
		ReporterLogSpans:            logSpans,
		ReporterLocalAgentHost:      reporterLocalAgentHost,
		ReporterLocalAgentPort:      reporterLocalAgentPort,
		ReporterCollectorEndpoint:   reporterCollectorEndpoint,
		ReporterCollectorUser:       reporterCollectorUser,
		ReporterCollectorPassword:   reporterCollectorPassword,
	}, nil
}
