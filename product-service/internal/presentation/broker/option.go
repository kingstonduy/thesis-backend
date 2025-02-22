package broker_server

import (
	"context"
	"encoding/json"

	cmd_pipeline "github.com/kingstonduy/go-core/comman-pipeline"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/trace"

	"github.com/kingstonduy/go-core/transport/broker"
	"github.com/kingstonduy/go-core/transport/broker/kafka"
	"github.com/kingstonduy/product-service/internal/domain"
	"golang.org/x/sync/errgroup"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"go.opentelemetry.io/otel"
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

		g.Go(func() error {
			return b.EventHandler(brokerCfg.OrderOutboxTopic)
		})

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
	logger.Infof(context.TODO(), "consume from topic=%s", topic)
	subscriber, err := b.Broker.Subscribe(topic, func(ctx context.Context, e broker.Event) (err error) {
		e.Ack() // auto ack

		var event domain.Event[cmd_pipeline.Outbox]
		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			logger.Errorf(ctx, "failed to unmarshal event: %v", err)
			return nil
		}

		ctx = otel.GetTextMapPropagator().Extract(context.Background(), otelsarama.NewConsumerMessageCarrier(&sarama.ConsumerMessage{
			Key:   []byte("traceparent"),
			Value: []byte(event.Payload.Before.TraceParent),
		}))

		outbox := event.Payload.After.ToOutboxWithTrace()
		ctx = trace.InjectTraceparent(ctx, outbox.TraceParent)
		err = b.dp.When(ctx, outbox)
		if err != nil {
			logger.Error(ctx, err)
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
