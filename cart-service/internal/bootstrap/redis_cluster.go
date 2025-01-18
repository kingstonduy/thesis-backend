package configuration

import (
	"context"

	"github.com/kingstonduy/go-core/logger"
	"github.com/redis/go-redis/v9"
)

func NewRedisClusterClient(cfg *Configuration) *redis.Client {
	redisConfig := cfg.RedisConfig

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addresses[0],
		Password: redisConfig.Password,
		// Addr:     redisConfig.Addr,
		// Password: redisConfig.Password,
	})

	// rdb := redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs:    redisConfig.Addresses,
	// 	Password: redisConfig.Password,
	// })

	logger.Infof(context.TODO(), "Conntected to redis cluster server: %v", redisConfig.Addresses)

	return rdb
}
