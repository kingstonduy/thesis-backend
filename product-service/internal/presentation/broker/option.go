package broker_server

import (
	"context"
	"encoding/json"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/go-core/transport/broker"
	"github.com/kingstonduy/go-core/transport/broker/kafka"
	"github.com/kingstonduy/product-service/internal/domain"
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
			return b.OutboxHandler(brokerCfg.OrderOutboxTopic)
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

func (b *BrokerServer) ProductCDCHandler(topic string) error {
	logger.Infof(context.TODO(), "consume from topic=%s", topic)
	subscriber, err := b.Broker.Subscribe(topic, func(c context.Context, e broker.Event) (err error) {
		e.Ack()

		var event domain.Event[domain.ProductCdc]
		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			logger.Errorf(context.TODO(), "failed to unmarshal event: %v", err)
			return err
		}

		logger.Infof(c, "consume message=%v", event)

		if err := b.redisClient.Set(c, "PRODUCT_ID"+"-"+event.Payload.After.ProductID, event.Payload.After, 0); err != nil {
			logger.Error(context.TODO(), "failed to set redis %v", err)
			return err
		}

		// logger.Infof(c, "received event: %v", event)

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

func (b *BrokerServer) OutboxHandler(topic string) error {
	ctx := context.TODO()
	logger.Infof(ctx, "consume from topic=%s", topic)
	subscriber, err := b.Broker.Subscribe(topic, func(c context.Context, e broker.Event) (err error) {
		e.Ack()

		var event domain.Event[domain.OutboxEntity]
		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			logger.Errorf(ctx, "failed to unmarshal event: %v", err)
			return err
		}

		switch event.Payload.After.CommandType {
		case domain.ORDER_INIT_TRANSACTION_COMMAND:
			cmd := domain.Command[domain.ExecuteTransactionRequest]{
				AggregateID: event.Payload.After.AggregateID,
				CommandID:   event.Payload.After.CommandID,
				CommandType: event.Payload.After.CommandType,
				ReplyTo:     event.Payload.After.ReplyTo,
			}

			if err := json.Unmarshal([]byte(event.Payload.After.Payload), &cmd.Payload); err != nil {
				logger.Errorf(ctx, "failed to unmarshal command: %v", err)
			} else {
				_, err = pipeline.Send[*domain.Command[domain.ExecuteTransactionRequest], *domain.ExecuteTransactionResponse](ctx, &cmd)
				if err != nil {
					logger.Error(ctx, err)
				}
			}
		}

		logger.Infof(c, "consume message=%v", event)
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
			logger.Error(ctx, err)
			return nil
		}

		_, err = pipeline.Send[*domain.Command[domain.RevertTransactionRequest], *domain.RevertTransactionResponse](ctx, &event.Payload.After)
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
