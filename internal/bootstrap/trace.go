package configuration

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/trace"
	"github.com/kingstonduy/go-core/trace/otel"
)

type Tr struct {
}

// ExtractSpanInfo implements trace.Tracer.
func (t *Tr) ExtractSpanInfo(context.Context) trace.SpanInfo {
	panic("unimplemented")
}

// StartTracing implements trace.Tracer.
func (t *Tr) StartTracing(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.SpanFinishFunc) {
	panic("unimplemented")
}

func NewTr() trace.Tracer {
	return &Tr{}
}

func GetTracer(cfg *Configuration) trace.Tracer {
	var (
		srvCfg   = cfg.ServerConfig
		traceCfg = cfg.TraceConfig
	)

	tracer, err := otel.NewOpenTelemetryTracer(
		context.Background(),
		trace.WithTraceServiceName(srvCfg.Name),
		trace.WithServiceVersion(srvCfg.AppVersion),
		trace.WithTraceExporterEndpoint(traceCfg.ExporterEndpoint),
	)

	if err != nil {
		panic(fmt.Errorf("failed to create tracer object: %w", err))
	}
	return tracer
}
