package trace

import (
	"io"

	otrace "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// NewTracer returns:
// - opentracing.Tracer for using in gin middleware
// io.closer for defering the tracer closer
// error on failure to create a new tracer
func NewTracer(cfg Config) (otrace.Tracer, io.Closer, error) {
	conf := cfg.New()
	tracer, closer, err := conf.New(
		cfg.App,
		config.Logger(jaeger.StdLogger),
	)

	if err != nil {
		return nil, nil, errCreateTracer
	}

	return tracer, closer, nil
}
