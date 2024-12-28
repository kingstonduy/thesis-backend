package main

import (
	"context"

	configuration "github.com/kingstonduy/comment-service/internal/bootstrap"
	"github.com/kingstonduy/comment-service/internal/postgres"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/server"

	http_server "github.com/kingstonduy/comment-service/internal/presentation/http"
	add_comment_uc "github.com/kingstonduy/comment-service/internal/usecase/add-comment"
	get_comment_uc "github.com/kingstonduy/comment-service/internal/usecase/get-comment"
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
	fx.Provide(add_comment_uc.NewAddCommentHandler),
	fx.Provide(get_comment_uc.NewGetCommentHandler),
)

var serverModule = fx.Module("server",
	fx.Provide(http_server.NewHttpServer),
)

var infraModule = fx.Module("infras",
	fx.Provide(postgres.NewProductRepoImpl),
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
