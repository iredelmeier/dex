package instrumentation

import (
	"os"

	lightstep "github.com/lightstep/lightstep-tracer-go"
	opentracing "github.com/opentracing/opentracing-go"
)

var (
	TracerAccessToken = os.Getenv("TRACER_ACCESS_TOKEN")
)

type TracerConfig struct {
	LightStep *lightstep.Options `json:"lightstep,omitempty=true"`
}

func NewTracer(c TracerConfig, componentName string) opentracing.Tracer {
	if ls := c.LightStep; ls != nil {
		ls.UseHttp = true

		if TracerAccessToken != "" {
			ls.AccessToken = TracerAccessToken
		}

		if host := os.Getenv("LIGHTSTEP_COLLECTOR_HOST"); host != "" {
			ls.Collector.Host = host
		}

		if ls.Tags == nil {
			ls.Tags = make(opentracing.Tags)
		}

		ls.Tags["component"] = componentName

		return lightstep.NewTracer(*ls)
	}

	return opentracing.NoopTracer{}
}
