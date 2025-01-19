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
	utils_transport "github.com/kingstonduy/order-service/internal/pkg/transport"
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
	tr := domain.TransactionEntity{
		TransactionID: trace.Sid,
		Status:        domain.INIT_STATUS,
		Processing:    1,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	var orderItems []domain.OrderEntity
	for _, item := range req.Details {
		orderItems = append(orderItems, domain.OrderEntity{
			OrderID:       uuid.New().String(),
			ProductID:     item.ProductID,
			UserID:        req.UserID,
			TransactionID: trace.Sid,
			CreatedAt:     now,
			UpdatedAt:     now,
		})
	}

	reqType := transport.Request[domain.ExecuteTransactionRequest]{
		Trace: utils_transport.GenRequestTrace(trace, "product-service", h.redisPubSub.GetChannel()),
		Data:  *req,
	}
	reqStr, _ := json.Marshal(reqType)
	outbox := domain.OutboxEntity{
		AggregateID: trace.Sid,
		CommandID:   uuid.New().String(),
		CommandType: domain.ORDER_INIT_TRANSACTION_COMMAND,
		Payload:     string(reqStr),
		ReplyTo:     h.redisPubSub.GetChannel(),
	}

	err2 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		err = h.transactionRepo.Insert(ctx, tr)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, item := range orderItems {
			err = h.orderRepo.Insert(ctx, item)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		err = h.outboxRepo.Insert(ctx, outbox)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}
		return nil
	})
	if err != nil || err2 != nil {
		if err == nil {
			err = err2
		}
		errx := errorx.FailedWithDetails(err.Error(), "")
		return nil, errx
	}

	resStr, err := h.redisPubSub.GetValue(ctx, trace.Sid, time.Second*20)
	if err != nil {
		errx := errorx.SuspendedErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx)
		return nil, errx
	}
	var resType transport.Response[domain.ExecuteTransactionResponse]
	err = json.Unmarshal([]byte(resStr), &resType)
	if err != nil {
		errx := errorx.SuspendedErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx)
		return nil, errx
	}

	if resType.Result.Code != transport.DefaultSuccessResponse.Result.Code {
		errx := errorx.SuspendedErrorWithDetails(resType.Result, "")
		logger.Error(ctx, errx)
		return nil, errx
	}

	return nil, nil
}
