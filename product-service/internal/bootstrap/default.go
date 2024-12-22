package configuration

import (
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/mapping"
	"github.com/kingstonduy/go-core/metrics"
	"github.com/kingstonduy/go-core/trace"
	"github.com/kingstonduy/go-core/validation"
)

func SetDefaults(
	log logger.Logger,
	tr trace.Tracer,
	ma mapping.Mapper,
	val validation.Validator,
	me *metrics.Metrics,
) {
	logger.SetDefaultLogger(log)
	trace.SetDefaultTracer(tr)
	mapping.SetDefaultMapper(ma)
	validation.SetDefaultValidator(val)
	metrics.SetDefaultMetrics(me)
}
