package main

import (
	"context"

	configuration "github.com/kingstonduy/authentication-service/internal/bootstrap"
	grpc_server "github.com/kingstonduy/authentication-service/internal/presentation"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/server"
	"go.uber.org/fx"
)

var configModule = fx.Module("config",
	fx.Provide(configuration.GetKafkaBroker),
	fx.Provide(configuration.NewCircuitBreaker),
	fx.Provide(configuration.GetConfigurationInstance),
	fx.Invoke(configuration.SetDefaults),
	fx.Provide(configuration.NewHealthChecker),
	fx.Provide(configuration.GetLogger),
	fx.Provide(configuration.GetMapper),
	fx.Provide(configuration.GetMetrics),
	fx.Provide(configuration.NewYugabyteCon),
	fx.Provide(configuration.NewRestyClient),
	fx.Provide(configuration.GetTracer),
	fx.Provide(configuration.GetValidator),
)

var serverModule = fx.Module("server",
	fx.Provide(grpc_server.NewGrpcServer),
)

func main() {
	fx.New(
		configModule,
		serverModule,
		fx.Invoke(run),
	).Run()
}

func run(
	lc fx.Lifecycle,
	grpcServer *grpc_server.GrpcServer,
	log logger.Logger,
	shutdowner fx.Shutdowner,
) {
	var gCtx = context.Background()
	var serverWrapper *server.ServerWrapper

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			serverWrapper = server.NewServerWrapper(
				server.WithStopHook(func(ctx context.Context) {
					log.Info(gCtx, "Stopped all the servers")
					shutdowner.Shutdown()
				}),
				server.WithServer("grpcServer", grpcServer),
			)

			return serverWrapper.Start(gCtx)
		},
		OnStop: func(ctx context.Context) error {
			if serverWrapper != nil {
				return serverWrapper.Stop(ctx)
			}
			return nil
		},
	})
}
