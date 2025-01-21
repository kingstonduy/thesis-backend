package main

import (
	"context"

	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/infra/postgres"
	add_cart_handler_uc "github.com/kingstonduy/cart-service/internal/usecase/add-cart-item"
	execute_transaction_uc "github.com/kingstonduy/cart-service/internal/usecase/execute-transaction"
	get_cart_uc "github.com/kingstonduy/cart-service/internal/usecase/get-cart"
	update_cart_uc "github.com/kingstonduy/cart-service/internal/usecase/update-cart-item"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/server"

	broker_server "github.com/kingstonduy/cart-service/internal/presentation/broker"
	http_server "github.com/kingstonduy/cart-service/internal/presentation/http"
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
	fx.Provide(add_cart_handler_uc.NewAddCartItemHandler),
	fx.Provide(get_cart_uc.NewGetCartHandler),
	fx.Provide(update_cart_uc.NewUpdateCartHandler),
	fx.Provide(execute_transaction_uc.NewExecuteTransactionHandler),
)

var serverModule = fx.Module("server",
	fx.Provide(http_server.NewHttpServer),
	fx.Provide(broker_server.NewBrokerServer),
)

var infraModule = fx.Module("infras",
	fx.Provide(postgres.NewCartRepo),
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
	HttpServer *http_server.HttpServer,
	BrokerServer *broker_server.BrokerServer,
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
				server.WithServer("BrokerServer", BrokerServer),
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
