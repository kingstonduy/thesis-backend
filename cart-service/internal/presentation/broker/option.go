package broker_server

import (
	"context"
	"encoding/json"

	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport/broker"
	"github.com/kingstonduy/go-core/transport/broker/kafka"
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

		// wait for the subscription result, return error if present
		if err := g.Wait(); err != nil {
			return err
		}

		brokerCfg := b.cfg.BrokerConfig

		b.CDCConsumer(brokerCfg.ProductCDCTopic)
		return nil
	}
}

func (b *BrokerServer) CDCConsumer(topic string) error {
	logger.Infof(context.TODO(), "consume from topic=%s", topic)
	subscriber, err := b.Broker.Subscribe(topic, func(c context.Context, e broker.Event) (err error) {
		e.Ack()

		var event domain.Event[domain.ProductCdc]
		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			logger.Errorf(c, "failed to unmarshal event: %v", err)
			return err
		}

		logger.Infof(c, "consume message=%v", event)

		if err := b.productRepo.Insert(c, event.Payload.After); err != nil {
			kMsg := broker.Message{
				Body:    e.Message().Body,
				Key:     e.Message().Key,
				Headers: e.Message().Headers,
			}
			if err1 := b.Broker.Publish(c, topic, &kMsg); err1 != nil {
				logger.Errorf(c, err1.Error())
			}
		}
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
