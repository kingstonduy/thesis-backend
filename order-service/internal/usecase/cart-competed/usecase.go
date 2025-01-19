package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/broker"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
)

type handler struct {
	orderRepo       domain.IOrderRepo
	transactionRepo domain.ITransactionRepo
	outboxRepo      domain.IOutboxRepo
	db              *configuration.PostgresCon
	kafka           broker.Broker
}

func NewCartCompltedHandler(
	orderRepo domain.IOrderRepo,
	transactionRepo domain.ITransactionRepo,
	outboxRepo domain.IOutboxRepo,
	db *configuration.PostgresCon,
	kafka broker.Broker,
) domain.ICartCompletedHandler {
	return &handler{
		orderRepo:       orderRepo,
		transactionRepo: transactionRepo,
		outboxRepo:      outboxRepo,
		db:              db,
		kafka:           kafka,
	}
}

// Handle implements domain.ICartCompletedHandler.
func (h *handler) Handle(ctx context.Context, cmd *domain.Command[transport.Request[domain.CartCompletedRequest]]) (res *domain.CartCompletedResponse, err error) {
	logger.Info(ctx, "CartCompletedHandler start")
	defer logger.Info(ctx, "CartCompletedHandler end")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC CartCompletedHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		cols := map[string]interface{}{
			"DELIVERY_STATUS": domain.COMPLETE_STATUS,
			"PAYMENT_STATUS":  domain.COMPLETE_STATUS,
		}
		conditions := map[string]interface{}{
			"TRANSACTION_ID": cmd.AggregateID,
		}

		if err = h.orderRepo.Update(ctx, cols, conditions); err != nil {
			logger.Error(ctx, err)
			return err
		}

		cols = map[string]interface{}{
			"STATUS":     domain.COMPLETE_STATUS,
			"PROCESSING": 0,
		}

		if err = h.transactionRepo.Update(ctx, cols, conditions); err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	})

	if err1 != nil || err != nil {
		if err == nil {
			err = err1
		}
		logger.Error(ctx, err)
		return nil, err
	}

	return nil, nil

}
