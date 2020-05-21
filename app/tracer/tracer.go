package tracer

import (
	"io"
	"log"

	"github.com/hashfyre/sample-go-app/app/config"
	"github.com/hashfyre/sample-go-app/pkg/trace"
	"github.com/opentracing/opentracing-go"
)

// InitTracer initializes a jaeger tracer from app config
func InitTracer() (opentracing.Tracer, io.Closer, error) {
	cfg, err := config.Get()
	if err != nil {
		return nil, nil, err
	}

	traceCfg := cfg.Tracer
	traceConfig := trace.Config{
		App: traceCfg.AppName,
		Sampler: trace.Sampler{
			Type: trace.ConstantSampler,
			Constant: trace.Constant{
				Enabled: true,
			},
		},
		Reporter: trace.Reporter{
			Type:      trace.LocalAgentReporter,
			QueueSize: traceCfg.ReporterQueueSize,
			LogSpans:  traceCfg.ReporterLogSpans,
			LocalAgent: trace.LocalAgent{
				Host: traceCfg.ReporterLocalAgentHost,
				Port: traceCfg.ReporterLocalAgentPort,
			},
		},
	}

	tracer, closer, err := trace.NewTracer(traceConfig)
	if err != nil {
		log.Println("Failed to initialize tracer")
		return nil, nil, err
	}

	return tracer, closer, nil
}
