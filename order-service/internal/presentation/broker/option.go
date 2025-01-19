package broker_server

import (
	"context"
	"encoding/json"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/broker"
	"github.com/kingstonduy/go-core/transport/broker/kafka"
	"github.com/kingstonduy/order-service/internal/domain"
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
			return b.EventHandler(brokerCfg.CartOutboxTopic)
		})

		// wait for the subscription result, return error if present
		if err := g.Wait(); err != nil {
			return err
		}

		return nil
	}
}

func (b *BrokerServer) EventHandler(topic string) error {
	ctx := context.TODO()
	logger.Infof(ctx, "consume from topic=%s", topic)
	subscriber, err := b.Broker.Subscribe(topic, func(c context.Context, e broker.Event) (err error) {
		e.Ack() // auto ack
		var event domain.Event[domain.OutboxEntity]
		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			logger.Errorf(ctx, "failed to unmarshal event: %v", err)
			return nil
		}

		switch event.Payload.After.CommandType {
		case domain.CART_COMPLETED_TRANSACTION_COMMAND:
			cmd := domain.Command[transport.Request[domain.CartCompletedRequest]]{
				AggregateID: event.Payload.After.AggregateID,
				CommandID:   event.Payload.After.CommandID,
				CommandType: event.Payload.After.CommandType,
				ReplyTo:     event.Payload.After.ReplyTo,
			}
			err = json.Unmarshal([]byte(event.Payload.After.Payload), &cmd.Payload)
			if err != nil {
				logger.Error(ctx, err)
				return nil
			}
			_, err := pipeline.Send[*domain.Command[transport.Request[domain.CartCompletedRequest]], *domain.CartCompletedResponse](ctx, &cmd)
			if err != nil {
				logger.Error(ctx, err)
			}

		default:
			logger.Errorf(ctx, "does not handle this event=%v", event)
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
