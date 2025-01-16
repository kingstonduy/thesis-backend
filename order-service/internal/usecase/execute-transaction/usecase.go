package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/broker"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
)

type handler struct {
	orderRepo       domain.IOrderRepo
	transactionRepo domain.ITransactionRepo
	db              *configuration.PostgresCon
	kafka           broker.Broker
}

func NewExecuteTransactionHandler(
	orderRepo domain.IOrderRepo,
	transactionRepo domain.ITransactionRepo,
	db *configuration.PostgresCon,
	kafka broker.Broker,
) domain.IExecuteTransactionHandler {
	return &handler{
		orderRepo:       orderRepo,
		transactionRepo: transactionRepo,
		db:              db,
		kafka:           kafka,
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
			ProductID:     reqDetails.ProductID,
			UserID:        req.UserID,
			TransactionID: transaction.TransactionID,
			CreatedAt:     now,
			UpdatedAt:     now,
		})
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

	// TODO add redis channel
	cmd := domain.Command[domain.ProductExecuteTransactionRequest]{
		AggregateID: trace.Sid,
		CommandID:   uuid.New().String(),
		CommandType: domain.EXECUTE_TRANSACTION_COMMAND,
		Payloay:     getProductExecuteTransactionRequest(*req),
		ReplyTo:     "redis-channel+uuid",
	}
	_ = cmd
	// h.kafka.Publish(ctx, "")

	// TODO use redis pubsub to wait

	return res, nil
}
