package main

import (
	"context"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/server"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/infra/postgres"
	execute_transaction_uc "github.com/kingstonduy/order-service/internal/usecase/execute-transaction"
	get_history_uc "github.com/kingstonduy/order-service/internal/usecase/get-history"

	http_server "github.com/kingstonduy/order-service/internal/presentation/http"
	redis_server "github.com/kingstonduy/order-service/internal/presentation/redis_pubsub"
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
	fx.Invoke(configuration.ResgisterPipeline),
	fx.Provide(configuration.NewYugabyteCon),
	fx.Provide(configuration.NewRedisClusterClient),
	fx.Provide(configuration.NewCacheClient),
	fx.Provide(configuration.NewRedixBroker),
	fx.Provide(configuration.NewRestyClient),
	fx.Provide(configuration.GetTracer),
	fx.Provide(configuration.GetValidator),
)

var usecaseModule = fx.Module("usecase",
	fx.Provide(execute_transaction_uc.NewExecuteTransactionHandler),
	fx.Provide(get_history_uc.NewGetHistoryHandler),
)

var serverModule = fx.Module("server",
	fx.Provide(http_server.NewHttpServer),
	fx.Provide(redis_server.NewRedisServer),
)

var infraModule = fx.Module("infras",
	fx.Provide(postgres.NewOrderRepo),
	fx.Provide(postgres.NewTransactionRepo),
	fx.Provide(postgres.NewOutboxRepo),
)

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
	redisServer *redis_server.RedisServer,
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
				server.WithServer("redisServer", redisServer),
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
