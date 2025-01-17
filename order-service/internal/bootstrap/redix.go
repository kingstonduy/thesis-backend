package configuration

import (
	redix "github.com/kingstonduy/order-service/internal/pkg/redis_broker"
	"github.com/redis/go-redis/v9"
)

func NewRedixBroker(
	redis *redis.Client,
) redix.PubSubBroker {
	return redix.NewRedisBroker(redis)
}
