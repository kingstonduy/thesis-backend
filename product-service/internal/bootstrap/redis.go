package configuration

import (
	"context"

	"github.com/kingstonduy/go-core/cache"
	"github.com/kingstonduy/go-core/cache/credis"
	"github.com/kingstonduy/go-core/logger"
	"github.com/redis/go-redis/v9"
)

func NewCacheClient(cfg *Configuration) cache.CacheClient {
	ctx := context.Background()
	redisConfig := cfg.RedisConfig

	client, err := credis.NewRedisClient(credis.WithRedisOptions(
		redis.UniversalOptions{
			Addrs:    redisConfig.Addresses,
			Password: redisConfig.Password,
		},
	))

	if err != nil {
		logger.Error(ctx, "Failted to connect to cache Client")
		panic(err)
	}

	logger.Infof(context.TODO(), "Conntected to redis server: %s", redisConfig.Addresses)

	return client
}
