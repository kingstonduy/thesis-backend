package configuration

import (
	"github.com/kingstonduy/go-core/metrics"
	"github.com/kingstonduy/go-core/metrics/prometheus"
)

var (
	KeyRequestTotal = []string{"request", "total"}
	KeyFailureTotal = []string{"failure", "total"}
)

func GetMetrics(cfg *Configuration) *metrics.Metrics {
	prome, err := prometheus.NewPrometheusSinkFrom(prometheus.PrometheusOpts{
		Expiration: 0,
		Name:       "prometheus_metrics",
	})

	if err != nil {
		panic(err)
	}

	config := metrics.DefaultConfig(cfg.ServerConfig.Name)
	config.EnableServiceLabel = true
	config.EnableHostname = false

	metrics, err := metrics.NewGlobal(config, prome)

	if err != nil {
		panic(err)
	}

	return metrics
}
