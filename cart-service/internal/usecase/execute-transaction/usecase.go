package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
	redix "github.com/kingstonduy/cart-service/internal/pkg/redis_broker"
	"github.com/kingstonduy/go-core/database"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/broker"
)

type handler struct {
	cfg         *configuration.Configuration
	cartRepo    domain.ICartRepo
	outboxRepo  domain.IOutboxRepo
	db          *configuration.PostgresCon
	redisPubSub redix.PubSubBroker
	kafka       broker.Broker
}

func NewExecuteTransactionHandler(
	cfg *configuration.Configuration,
	cartRepo domain.ICartRepo,
	outboxRepo domain.IOutboxRepo,
	db *configuration.PostgresCon,
	redisPubSub redix.PubSubBroker,
	kafka broker.Broker,
) domain.IExecuteTransactionHandler {
	return &handler{
		cfg:         cfg,
		cartRepo:    cartRepo,
		outboxRepo:  outboxRepo,
		db:          db,
		redisPubSub: redisPubSub,
		kafka:       kafka,
	}
}

// Handle implements domain.IExecuteTransactionHandler.
func (h *handler) Handle(ctx context.Context, cmd *domain.Command[domain.ExecuteTransactionRequest]) (res *domain.ExecuteTransactionResponse, err error) {
	logger.Info(ctx, "Get products handler start")
	defer logger.Info(ctx, "Get products handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	req := cmd.Payload

	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		err = fmt.Errorf("error simulating")
		return nil

		for _, item := range req.Details {
			err = h.cartRepo.DeleteById(ctx, item.CartItemID)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		result := transport.DefaultSuccessResponse.Result
		resultStr, _ := json.Marshal(result)
		outbox := domain.OutboxEntity{
			AggregateID: cmd.AggregateID,
			CommandID:   uuid.New().String(),
			CommandType: domain.CART_COMPLETED_TRANSACTION_COMMAND,
			Payload:     string(resultStr),
			ReplyTo:     cmd.ReplyTo,
		}
		err = h.outboxRepo.Insert(ctx, outbox)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	}, database.WithIsolationLevelOptions(sql.LevelReadUncommitted))

	if err1 != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())

		event := domain.Event[domain.Command[domain.ExecuteTransactionRequest]]{
			Payload: domain.EventPayload[domain.Command[domain.ExecuteTransactionRequest]]{
				After: domain.Command[domain.ExecuteTransactionRequest]{
					AggregateID: cmd.AggregateID,
					CommandID:   uuid.New().String(),
					CommandType: domain.CART_FAILED_TRANSACTION_COMMAND,
					Payload:     cmd.Payload,
					ReplyTo:     cmd.ReplyTo,
				},
			},
		}
		eventStr, _ := json.Marshal(event)
		kMsg := broker.Message{
			Key:  []byte(cmd.AggregateID),
			Body: eventStr,
		}
		if err = h.kafka.Publish(ctx, h.cfg.BrokerConfig.CartTopic, &kMsg); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
		return nil, errx
	}

	if err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())

		event := domain.Event[domain.Command[domain.ExecuteTransactionRequest]]{
			Payload: domain.EventPayload[domain.Command[domain.ExecuteTransactionRequest]]{
				After: domain.Command[domain.ExecuteTransactionRequest]{
					AggregateID: cmd.AggregateID,
					CommandID:   uuid.New().String(),
					CommandType: domain.CART_FAILED_TRANSACTION_COMMAND,
					Payload:     cmd.Payload,
					ReplyTo:     cmd.ReplyTo,
				},
			},
		}
		eventStr, _ := json.Marshal(event)
		kMsg := broker.Message{
			Key:  []byte(cmd.AggregateID),
			Body: eventStr,
		}
		if err = h.kafka.Publish(ctx, h.cfg.BrokerConfig.CartTopic, &kMsg); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
		return nil, errx
	}

	return res, nil
}
