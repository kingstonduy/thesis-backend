package configuration

import (
	"context"
	"time"

	"github.com/kingstonduy/go-core/logger"
	"github.com/sony/gobreaker"
)

// shouldBeSwitchedToOpen checks if the circuit breaker should
// switch to the Open state
func shouldBeSwitchedToOpen(counts gobreaker.Counts) bool {
	failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
	return counts.Requests >= 3 && failureRatio >= 0.6
}

func NewCircuitBreaker(logger logger.Logger) *gobreaker.CircuitBreaker {
	cfg := gobreaker.Settings{
		// When to flush counters int the Closed state
		Interval: 5 * time.Second,
		// Time to switch from Open to Half-open
		Timeout: 7 * time.Second,
		// Function with check when to switch from Closed to Open
		ReadyToTrip: shouldBeSwitchedToOpen,
		OnStateChange: func(_ string, from gobreaker.State, to gobreaker.State) {
			// Handler for every state change. We'll use for debugging purpose
			logger.Infof(context.Background(), "state changed from %v to %v", from.String(), to.String())
		},
	}

	return gobreaker.NewCircuitBreaker(cfg)
}
