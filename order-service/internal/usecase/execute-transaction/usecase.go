package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/broker"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
	redix "github.com/kingstonduy/order-service/internal/pkg/redis_broker"
)

type handler struct {
	orderRepo       domain.IOrderRepo
	transactionRepo domain.ITransactionRepo
	outboxRepo      domain.IOutboxRepo
	db              *configuration.PostgresCon
	kafka           broker.Broker
	redisPubSub     redix.PubSubBroker
}

func NewExecuteTransactionHandler(
	orderRepo domain.IOrderRepo,
	transactionRepo domain.ITransactionRepo,
	outboxRepo domain.IOutboxRepo,
	db *configuration.PostgresCon,
	kafka broker.Broker,
	redisPubSub redix.PubSubBroker,
) domain.IExecuteTransactionHandler {
	return &handler{
		orderRepo:       orderRepo,
		transactionRepo: transactionRepo,
		outboxRepo:      outboxRepo,
		db:              db,
		kafka:           kafka,
		redisPubSub:     redisPubSub,
	}
}

// Handle implements domain.IExecuteTransactionHandler.
func (h *handler) Handle(ctx context.Context, req *domain.ExecuteTransactionRequest) (res *domain.ExecuteTransactionResponse, err error) {
	logger.Info(ctx, "ExecuteTransactionHandler start")
	defer logger.Info(ctx, "ExecuteTransactionHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC ExecuteTransactionHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	trace := transport.GetTraceByCtx(ctx)

	now := time.Now()
	transaction := domain.TransactionEntity{
		TransactionID: uuid.New().String(),
		Status:        domain.INIT_STATUS,
		Processing:    1,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	var orderItems = []domain.OrderEntity{}
	for _, reqDetails := range req.Details {
		orderItems = append(orderItems, domain.OrderEntity{
			OrderID:       uuid.New().String(),
			ProductID:     reqDetails.ProductID,
			UserID:        req.UserID,
			TransactionID: transaction.TransactionID,
			CreatedAt:     now,
			UpdatedAt:     now,
		})
	}

	payloadStr, _ := json.Marshal(req)
	var outbox domain.OutboxEntity = domain.OutboxEntity{
		AggregateID: trace.Sid,
		CommandID:   uuid.New().String(),
		CommandType: domain.ORDER_INIT_TRANSACTION_COMMAND,
		Payloay:     string(payloadStr),
		ReplyTo:     h.redisPubSub.GetChannel(),
	}

	// insert db
	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, item := range orderItems {
			err = h.orderRepo.Insert(ctx, item)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		err = h.transactionRepo.Insert(ctx, transaction)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		err = h.outboxRepo.Insert(ctx, outbox)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	})
	if err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}
	if err1 != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	// TODO add redis pubsub here to listen event
	resultStr, err := h.redisPubSub.GetValue(ctx, trace.Sid, time.Second*10)
	if err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	var result transport.Result
	err = json.Unmarshal([]byte(resultStr), &result)
	if err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	if result.Code == errorx.DefaultSuccessResponseCode {
		return nil, nil
	} else {
		errx := &errorx.Error{
			Code:    result.Code,
			Status:  result.StatusCode,
			Message: result.Message,
			Details: result.Details,
		}
		logger.Error(ctx, errx.Error())
		return nil, errx
	}
}
