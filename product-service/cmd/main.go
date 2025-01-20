package main

import (
	"context"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/server"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/infra/postgres"
	execute_transaction_uc "github.com/kingstonduy/product-service/internal/usecase/execute-transaction"
	get_products_uc "github.com/kingstonduy/product-service/internal/usecase/get-all-product"
	get_product_by_category_uc "github.com/kingstonduy/product-service/internal/usecase/get-product-by-category"
	get_product_by_gender_uc "github.com/kingstonduy/product-service/internal/usecase/get-product-by-gender"
	get_product_detail_uc "github.com/kingstonduy/product-service/internal/usecase/get-product-detail"
	revert_transaction_uc "github.com/kingstonduy/product-service/internal/usecase/revert-transaction"

	broker_server "github.com/kingstonduy/product-service/internal/presentation/broker"
	http_server "github.com/kingstonduy/product-service/internal/presentation/http"
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
	fx.Provide(get_products_uc.NewGetProductsHandler),
	fx.Provide(get_product_detail_uc.NewGetProductDetailHandler),
	fx.Provide(execute_transaction_uc.NewExecuteTransactionHandler),
	fx.Provide(revert_transaction_uc.NewRevertTransactionHandler),
	fx.Provide(get_product_by_gender_uc.NewGetProductsByGenderHandler),
	fx.Provide(get_product_by_category_uc.NewGetProductsByCategoryHandler),
)

var serverModule = fx.Module("server",
	fx.Provide(http_server.NewHttpServer),
	fx.Provide(broker_server.NewBrokerServer),
)

var infraModule = fx.Module("infras",
	fx.Provide(postgres.NewProductRepoImpl),
	fx.Provide(postgres.NewOutboxRepo),
	fx.Provide(postgres.NewInventoryRepoImpl),
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
	brokerServer *broker_server.BrokerServer,
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
				server.WithServer("brokerServer", brokerServer),
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
