package broker_server

import (
	"context"
	"encoding/json"

	"github.com/kingstonduy/cart-service/internal/domain"
	redix "github.com/kingstonduy/cart-service/internal/pkg/redis_broker"
	utils_transport "github.com/kingstonduy/cart-service/internal/pkg/transport"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/go-core/transport"
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

		brokerCfg := b.cfg.BrokerConfig
		// TODO open later
		// g.Go(func() error {
		// 	return b.ProductCDCHandler(brokerCfg.ProductCDCTopic)
		// })
		g.Go(func() error {
			return b.EventHandler(brokerCfg.ProductOutboxTopic)
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
		case domain.PRODUCT_COMPLETED_TRANSACTION_COMMAND:
			cmd := domain.Command[transport.Request[domain.ExecuteTransactionRequest]]{
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
			_, err := pipeline.Send[*domain.Command[transport.Request[domain.ExecuteTransactionRequest]], *domain.ExecuteTransactionResponse](ctx, &cmd)
			if err != nil {
				logger.Error(ctx, err)
				// cast err to errorx
				errx, ok := err.(*errorx.Error)
				if !ok {
					errx = errorx.InternalServerErrorWithDetails(err.Error(), "")
				}

				resType := transport.Response[any]{
					Trace:  utils_transport.GenResponseTrace(cmd.Payload.Trace),
					Result: *utils_transport.GenResultFromErrorx(ctx, errx),
				}
				err := b.redisPubSub.Publish(ctx, redix.NewMessage(cmd.AggregateID, resType, cmd.ReplyTo))
				if err != nil {
					logger.Error(ctx, err)
					return nil
				}
				return nil
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
