package broker_server

import (
	"context"
	"encoding/json"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/broker"
	"github.com/kingstonduy/go-core/transport/broker/kafka"
	"github.com/kingstonduy/order-service/internal/domain"
	redix "github.com/kingstonduy/order-service/internal/pkg/redis_broker"
	"golang.org/x/sync/errgroup"
)

func (b *BrokerServer) GetSubscriptionOptions() []broker.SubscribeOption {
	return []broker.SubscribeOption{
		broker.WithSubscribeGroup(b.consumerGroup),
		kafka.InitialOffset(-1), // latest offset
		broker.WithSubscribeAutoAck(true),
	}
}

func WithSubscriptions() BrokerServerStartOption {
	return func(b *BrokerServer) error {
		g := new(errgroup.Group)

		brokerCfg := b.cfg.BrokerConfig
		// TODO open later
		// g.Go(func() error {
		// 	return b.ProductCDCHandler(brokerCfg.ProductCDCTopic)
		// })

		g.Go(func() error {
			return b.CartOutboxHandler(brokerCfg.CartOutboxTopic)
		})

		g.Go(func() error {
			return b.CartFailedHandler(brokerCfg.CartTopic)
		})

		// wait for the subscription result, return error if present
		if err := g.Wait(); err != nil {
			return err
		}

		return nil
	}
}

func (b *BrokerServer) CartOutboxHandler(topic string) error {
	ctx := context.TODO()
	logger.Infof(ctx, "consume from topic=%s", topic)
	subscriber, err := b.Broker.Subscribe(topic, func(c context.Context, e broker.Event) (err error) {
		e.Ack() // auto ack
		var event domain.Event[domain.OutboxEntity]
		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			logger.Errorf(ctx, "failed to unmarshal event: %v", err)
			return nil
		}

		err = b.redisPubSub.Publish(ctx, redix.RedisMessage{
			Key:     event.Payload.After.AggregateID,
			Value:   event.Payload.After.Payload,
			Channel: event.Payload.After.ReplyTo,
		})
		if err != nil {
			logger.Error(ctx, err)
			return nil
		}

		logger.Infof(c, "consumed message=%v", event)
		return nil
	}, b.GetSubscriptionOptions()...)
	if err != nil {
		return err
	}
	go func() {
		defer subscriber.Unsubscribe()
		<-b.quit
	}()
	return nil
}

func (b *BrokerServer) CartFailedHandler(topic string) error {
	ctx := context.TODO()
	logger.Infof(ctx, "consume from topic=%s", topic)
	subscriber, err := b.Broker.Subscribe(topic, func(c context.Context, e broker.Event) (err error) {
		e.Ack() // auto ack
		var event domain.Event[domain.Command[domain.RevertTransactionRequest]]
		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			logger.Errorf(ctx, "failed to unmarshal event: %v", err)
			return nil
		}

		result := transport.DefaultFailureResponse.Result
		resultStr, _ := json.Marshal(result)

		err = b.redisPubSub.Publish(ctx, redix.RedisMessage{
			Key:     event.Payload.After.AggregateID,
			Value:   string(resultStr),
			Channel: event.Payload.After.ReplyTo,
		})
		if err != nil {
			logger.Error(ctx, err)
			return nil
		}

		logger.Infof(c, "consumed message=%v", event)
		return nil
	}, b.GetSubscriptionOptions()...)
	if err != nil {
		return err
	}
	go func() {
		defer subscriber.Unsubscribe()
		<-b.quit
	}()
	return nil
}
