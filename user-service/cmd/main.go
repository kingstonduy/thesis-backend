package main

import (
	"context"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/server"
	configuration "github.com/kingstonduy/user-service/internal/bootstrap"
	"github.com/kingstonduy/user-service/internal/infra/postgres"
	get_user_info_uc "github.com/kingstonduy/user-service/internal/usecase/get-user-information"
	login_uc "github.com/kingstonduy/user-service/internal/usecase/login"
	register_uc "github.com/kingstonduy/user-service/internal/usecase/register"
	update_user_uc "github.com/kingstonduy/user-service/internal/usecase/update-user-information"

	http_server "github.com/kingstonduy/user-service/internal/presentation/http"
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
	fx.Invoke(configuration.RegisterPipeline),
	fx.Provide(configuration.NewYugabyteCon),
	fx.Provide(configuration.NewRestyClient),
	fx.Provide(configuration.GetTracer),
	fx.Provide(configuration.GetValidator),
)

var usecaseModule = fx.Module("usecase",
	fx.Provide(get_user_info_uc.NewGetUserInformationhandler),
	fx.Provide(login_uc.NewLoginHandler),
	fx.Provide(register_uc.NewRegisterHandler),
	fx.Provide(update_user_uc.NewUpdateHandler),
)

var serverModule = fx.Module("server",
	fx.Provide(http_server.NewHttpServer),
)

var infraModule = fx.Module("infras",
	fx.Provide(postgres.NewCustomerRepo),
	fx.Provide(postgres.NewLocationRepo),
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
