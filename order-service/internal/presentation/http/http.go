package http_server

import (
	"context"
	"fmt"

	healthchecks "github.com/kingstonduy/go-core/health"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/trace"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
	"github.com/kingstonduy/go-core/validation"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/pkg/errors"
)

type HttpServer struct {
	Port         int
	Name         string
	App          *fiberx.FiberApp
	HealhChecker healthchecks.HealthChecker
	Validator    validation.Validator
	Tracer       trace.Tracer
	connected    bool
}

func NewHttpServer(
	cfg *configuration.Configuration,
	logger logger.Logger,
	healhChecker healthchecks.HealthChecker,
	validator validation.Validator,
	tracer trace.Tracer,
) *HttpServer {
	fiberApp := fiberx.NewFiberApp(
		fiberx.WithLogger(logger),
		fiberx.WithTracer(tracer),
		fiberx.WithBasePath(cfg.ServerConfig.HttpBasePath),
		fiberx.WithSwaggerPath("/swagger/*"),
		fiberx.WithServiceName(cfg.ServerConfig.Name),
		fiberx.WithRateLimiterEnabled(false),
	)

	return &HttpServer{
		Port:         cfg.ServerConfig.HttpPort,
		Name:         cfg.ServerConfig.Name,
		App:          fiberApp,
		HealhChecker: healhChecker,
		Validator:    validator,
		Tracer:       tracer,
	}
}

func (s *HttpServer) getAllOptions() []option {
	return []option{
		s.WithRoutingOption(),
	}
}

func (s *HttpServer) Start(ctx context.Context) error {
	opts := s.getAllOptions()
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return err
		}
	}

	logger.Infof(ctx, "Start HTTP server at port: %v", s.Port)

	s.connected = true
	if err := s.App.Listen(fmt.Sprintf(":%v", s.Port)); err != nil {
		s.connected = false
		err = errors.Wrap(err, "failed to start http server")
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	if err := s.App.Shutdown(); err != nil {
		err = errors.Wrap(err, "failed to shutdown http server")
		logger.Error(ctx, err)
		return err
	}

	s.connected = false
	logger.Info(ctx, "Stopped HTTP Server")
	return nil
}

// Connected implements presentation.Server.
func (s *HttpServer) Connected() bool {
	return s.connected
}

type option func(s *HttpServer) error
