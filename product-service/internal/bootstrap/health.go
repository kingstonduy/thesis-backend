package configuration

import (
	"time"

	health "github.com/kingstonduy/go-core/health"
)

func NewHealthChecker(cfg *Configuration) health.HealthChecker {
	var (
		hConfig = cfg.HealthCheckConfig
		sConfig = cfg.ServerConfig
	)

	// Init health
	healthChecker := health.NewHealthChecker(
		health.WithName(sConfig.Name),
		health.WithVersion(sConfig.AppVersion),
	)

	// check Garbage Collector
	gcChecker := health.NewGCMaxChecker(time.Millisecond * time.Duration(hConfig.GcMaxPauseThresholdms))
	healthChecker.AddLivenessCheck("garbage collector check", gcChecker)

	// check Goroutine
	grChecker := health.NewGoroutineChecker(hConfig.GrRunningThreshold)
	healthChecker.AddLivenessCheck("goroutine checker", grChecker)

	// check network
	pingChecker := health.NewPingChecker("http://google.com", "GET", time.Millisecond*time.Duration(200), nil, nil)
	healthChecker.AddReadinessCheck("ping check", pingChecker)

	return healthChecker
}
