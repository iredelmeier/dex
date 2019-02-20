package instrumentation

import (
	opentracing "github.com/opentracing/opentracing-go"
)

func NewTracer() opentracing.Tracer {
	return opentracing.NoopTracer{}
}
