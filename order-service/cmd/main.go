package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/server"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/infra/postgres"
	cart_completed_uc "github.com/kingstonduy/order-service/internal/usecase/cart-competed"
	execute_transaction_uc "github.com/kingstonduy/order-service/internal/usecase/execute-transaction"
	get_history_uc "github.com/kingstonduy/order-service/internal/usecase/get-history"
	revert_transaction_uc "github.com/kingstonduy/order-service/internal/usecase/revert-transaction"

	broker_server "github.com/kingstonduy/order-service/internal/presentation/broker"
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
	fx.Provide(configuration.NewDispatcher),
)

var usecaseModule = fx.Module("usecase",
	fx.Provide(execute_transaction_uc.NewExecuteTransactionHandler),
	fx.Provide(get_history_uc.NewGetHistoryHandler),
	fx.Provide(cart_completed_uc.NewCartCompltedHandler),
	fx.Provide(revert_transaction_uc.NewRevertTransactionHandler),
)

var serverModule = fx.Module("server",
	fx.Provide(http_server.NewHttpServer),
	fx.Provide(redis_server.NewRedisServer),
	fx.Provide(broker_server.NewBrokerServer),
)

var infraModule = fx.Module("infras",
	fx.Provide(postgres.NewOrderRepo),
	fx.Provide(postgres.NewTransactionRepo),
	fx.Provide(postgres.NewOutboxRepo),
	fx.Provide(postgres.NewProductRepo),
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
	BrokerServer *broker_server.BrokerServer,
	log logger.Logger,
	shutdowner fx.Shutdowner,
) {
	var gCtx = context.Background()
	var serverWrapper *server.ServerWrapper

	// Load the location for GMT+7
	_, err := time.LoadLocation("Asia/Bangkok") // Bangkok is in GMT+7
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Get the current time in the specified location
	currentTime := time.Now()
	fmt.Println("💡💡💡💡💡💡💡💡💡" + currentTime.Format("2006-01-02T15:04:05"))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			serverWrapper = server.NewServerWrapper(
				server.WithStopHook(func(ctx context.Context) {
					log.Info(gCtx, "Stopped all the servers")
					shutdowner.Shutdown()
				}),
				server.WithServer("httpServer", HttpServer),
				server.WithServer("redisServer", redisServer),
				server.WithServer("brokerServer", BrokerServer),
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
