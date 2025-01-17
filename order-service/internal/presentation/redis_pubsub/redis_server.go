package redis_server

import (
	"context"

	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	redix "github.com/kingstonduy/order-service/internal/pkg/redis_broker"
	"github.com/pkg/errors"
)

type RedisServer struct {
	cfg         *configuration.Configuration
	redisBroker redix.PubSubBroker
	connected   bool
}

func NewRedisServer(
	cfg *configuration.Configuration,
	redisBroker redix.PubSubBroker,
) *RedisServer {
	RedisServer := &RedisServer{
		cfg:         cfg,
		redisBroker: redisBroker,
	}

	return RedisServer
}

func (s *RedisServer) Start(ctx context.Context) error {
	// opts := s.getAllOptions()
	// for _, opt := range opts {
	// 	if err := opt(s); err != nil {
	// 		return err
	// 	}
	// }

	logger.Infof(ctx, "Start redis server successfully on channel=%s", s.redisBroker.GetChannel())

	s.connected = true
	if err := s.redisBroker.Listen(ctx); err != nil {
		s.connected = false
		err = errors.Wrap(err, "failed to start to redis broker")
		return err
	}

	return nil
}

func (s *RedisServer) Connected() bool {
	return s.connected
}

func (s *RedisServer) Stop(ctx context.Context) error {
	if err := s.redisBroker.Shutdown(context.Background()); err != nil {
		err = errors.Wrap(err, "failed to shutdown redis server")
		logger.Error(ctx, err)
		return err
	}

	s.connected = false

	logger.Info(ctx, "stopped redis server")
	return nil
}
