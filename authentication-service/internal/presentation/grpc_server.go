package grpc_server

import (
	"context"
	"fmt"
	"net"
	"time"

	configuration "github.com/kingstonduy/authentication-service/internal/bootstrap"
	authentication_proto "github.com/kingstonduy/authentication-service/proto"
	healthchecks "github.com/kingstonduy/go-core/health"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/trace"
	"github.com/kingstonduy/go-core/validation"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Port         int
	Name         string
	App          *grpc.Server
	HealhChecker healthchecks.HealthChecker
	Validator    validation.Validator
	Tracer       trace.Tracer
	connected    bool
}

func NewGrpcServer(
	cfg *configuration.Configuration,
	logger logger.Logger,
	healhChecker healthchecks.HealthChecker,
	validator validation.Validator,
	tracer trace.Tracer,
) *GrpcServer {
	app := grpc.NewServer(grpc.ConnectionTimeout(time.Second * 60))
	authentication_proto.RegisterAuthenticationServiceServer(app, NewApp())

	return &GrpcServer{
		Port:         cfg.ServerConfig.GrpcPort,
		Name:         cfg.ServerConfig.Name,
		App:          app,
		HealhChecker: healhChecker,
		Validator:    validator,
		Tracer:       tracer,
	}
}

func (s *GrpcServer) getAllOptions() []option {
	return []option{}
}

func (s *GrpcServer) Start(ctx context.Context) error {
	opts := s.getAllOptions()
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return err
		}
	}

	logger.Infof(ctx, "Start Grpc server at port: %v", s.Port)

	s.connected = true

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.Port))
	if err != nil {
		s.connected = false
		err = errors.Wrap(err, "failed to start Grpc server")
		logger.Error(ctx, err)
		return err
	}

	if err = s.App.Serve(lis); err != nil {
		s.connected = false
		err = errors.Wrap(err, "failed to start Grpc server")
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (s *GrpcServer) Stop(ctx context.Context) error {
	s.App.GracefulStop()

	s.connected = false
	logger.Info(ctx, "Stopped Grpc Server")
	return nil
}

// Connected implements presentation.Server.
func (s *GrpcServer) Connected() bool {
	return s.connected
}

type option func(s *GrpcServer) error
