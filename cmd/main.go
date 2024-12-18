package main

import (
	"context"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/server"
	configuration "github.com/kingstonduy/thesis-backend/internal/bootstrap"

	http_server "github.com/kingstonduy/thesis-backend/internal/presentation/http"
	"go.uber.org/fx"
)

var configModule = fx.Module("config",
	fx.Provide(configuration.GetLogger),
	fx.Provide(configuration.GetConfigure),
	fx.Provide(configuration.GetConfigurationInstance),
	fx.Provide(configuration.GetValidator),
	fx.Provide(configuration.GetTracer),
	fx.Provide(configuration.GetKafkaBroker),
	fx.Provide(configuration.NewHealthChecker),
	fx.Provide(configuration.GetMetrics),
	fx.Invoke(configuration.SetDefaults),
	fx.Provide(configuration.GetMapper),
	fx.Provide(configuration.NewCircuitBreaker),
	fx.Provide(configuration.NewRestyClient),
	fx.Invoke(configuration.ResgisterPipeline),
	fx.Provide(configuration.GetYugabyteMcsAssetMgmtDataCon),
	fx.Provide(configuration.GetOracleOsbr20Con),
	fx.Provide(configuration.GetOracleUatsanCon),
)

var usecaseModule = fx.Module("usecase")
var serverModule = fx.Module("server",
	fx.Provide(http_server.NewHttpServer),
)

var infraModule = fx.Module("infras")

func main() {
	fx.New(
		configModule,
		usecaseModule,
		serverModule,
		infraModule,
		fx.Invoke(run),
	).Run()
}

func run(
	lc fx.Lifecycle,
	HttpServer *http_server.HttpServer,
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
				server.WithServer("httpServer", HttpServer),
				// server.WithServer("cronjobServer", cronjobServer),
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
