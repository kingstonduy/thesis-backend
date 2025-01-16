package redix

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type RedisMessage struct {
	Key     string `json:"key" validate:"required"`
	Value   string `json:"value" validate:"required"`
	Channel string `json:"channel" validate:"required"`
}

type PubSubBroker interface {
	Listen(ctx context.Context) error
	Publish(ctx context.Context, rMsg RedisMessage) error
	Shutdown(ctx context.Context) error
	GetValue(ctx context.Context, key string, dur time.Duration) (value string, err error)
	GetChannel() string
}

// redisBroker holds the map of channels and a Redis client
type redisBroker struct {
	chans cmap.ConcurrentMap[string, chan string]
	// redis       *redis.Client
	redis       *redis.ClusterClient
	channel     string
	subscribers *redis.PubSub
	validator   *validator.Validate
}

type option func(*redisBroker)

// Set the topic which redis consume
func WithChannel(channel string) option {
	return func(c *redisBroker) {
		c.channel = channel
	}
}

func NewRedisBroker(
	redis *redis.ClusterClient,
	opts ...option,
) PubSubBroker {
	chans := cmap.New[chan string]()
	channel := uuid.New().String()

	manager := &redisBroker{
		chans:     chans,
		redis:     redis,
		channel:   channel,
		validator: validator.New(),
	}

	// Apply all options
	for _, opt := range opts {
		opt(manager)
	}

	// Subscribe to the updated channel (if the channel was changed by an option)
	manager.subscribers = redis.Subscribe(context.Background(), manager.channel)

	return manager
}

// Listen implements PubSubBroker.
func (c *redisBroker) Listen(ctx context.Context) error {
	pubsub := c.redis.Subscribe(ctx, c.channel)
	// Close the subscription when we are done.
	defer pubsub.Close()
	ch := pubsub.Channel()
	for msg := range ch {
		c.handle(ctx, msg)
	}
	return nil
}

func (c *redisBroker) handle(ctx context.Context, msg *redis.Message) {
	var rMsg RedisMessage
	if err := json.Unmarshal([]byte(msg.Payload), &rMsg); err != nil {
	} else {
		if !c.chans.Has(rMsg.Key) {
			return
		}

		ch, _ := c.chans.Get(rMsg.Key)

		// Send the message to the channel in a non-blocking way
		ch <- rMsg.Value // Assuming command.Data is of type string
	}
}

// Publish implements PubSubBroker.
func (c *redisBroker) Publish(ctx context.Context, redisMsg RedisMessage) error {
	if err := c.validator.Struct(redisMsg); err != nil {
		return err
	}

	// Ensure the channel for the key exists before publishing
	_, exist := c.chans.Get(redisMsg.Key)
	if !exist {
		ch := make(chan string, 5)
		c.chans.Set(redisMsg.Key, ch)
	}

	// Publish to Redis
	redisBytes, _ := json.Marshal(redisMsg)
	if err := c.redis.Publish(ctx, redisMsg.Channel, string(redisBytes)).Err(); err != nil {
		return errors.Wrap(err, "can not publish redis")
	}

	return nil
}

// Shutdown implements PubSubBroker.
func (c *redisBroker) Shutdown(ctx context.Context) error {
	// Gracefully unsubscribe from Redis channels
	if err := c.subscribers.Unsubscribe(ctx, c.channel); err != nil {
		return err
	}

	// Close the PubSub connection
	if err := c.subscribers.Close(); err != nil {
		return err
	}

	return nil
}

func (c *redisBroker) GetValue(ctx context.Context, key string, dur time.Duration) (value string, err error) {

	// Check if the channel exists, if not create it
	ch, exist := c.chans.Get(key)
	if !exist {
		ch = make(chan string, 5) // Create a new buffered channel
		c.chans.Set(key, ch)      // Save the channel in the map
	}

	// Wait for a value or timeout
	select {
	case value = <-ch:
		c.chans.Remove(key) // Remove the channel after receiving the value
		return value, nil
	case <-time.After(dur):
		c.chans.Remove(key) // Remove the channel after receiving the value
		return "", fmt.Errorf("timeout waiting for value from key '%s'", key)
	}
}

func (c *redisBroker) GetChannel() string {
	return c.channel
}
