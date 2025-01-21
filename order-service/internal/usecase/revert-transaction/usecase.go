package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
)

type handler struct {
	orderRepo       domain.IOrderRepo
	transactionRepo domain.ITransactionRepo
	db              *configuration.PostgresCon
}

func NewRevertTransactionHandler(
	orderRepo domain.IOrderRepo,
	transactionRepo domain.ITransactionRepo,
	db *configuration.PostgresCon,
) domain.IRevertTransactionHandler {
	return &handler{
		orderRepo:       orderRepo,
		transactionRepo: transactionRepo,
		db:              db,
	}
}

// Handle implements domain.IRevertTransactionHandler.
func (h *handler) Handle(ctx context.Context, cmd *domain.Command[transport.Request[domain.RevertTransactionRequest]]) (res *domain.RevertTransactionResponse, err error) {
	logger.Info(ctx, "RevertTransaction handler start")
	defer logger.Info(ctx, "RevertTransaction handler end")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC RevertTransaction handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		tr, err := h.transactionRepo.SelectByTransactionID(ctx, cmd.AggregateID)
		if err != nil {
			return err
		}

		cols := map[string]interface{}{
			"STATUS":     domain.FAILED_STATUS,
			"PROCESSING": 0,
			"UPDATED_AT": time.Now(),
		}
		conditions := map[string]interface{}{
			"TRANSACTION_ID": tr.TransactionID,
			"UPDATED_AT":     tr.UpdatedAt,
		}
		err = h.transactionRepo.Update(ctx, cols, conditions)
		if err != nil {
			return err
		}

		cols = map[string]interface{}{
			"DELIVERY_STATUS": domain.CANCEL_STATUS,
			"PAYMENT_STATUS":  domain.CANCEL_STATUS,
			"UPDATED_AT":      time.Now(),
		}
		conditions = map[string]interface{}{
			"TRANSACTION_ID": tr.TransactionID,
		}
		err = h.orderRepo.Update(ctx, cols, conditions)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil || err1 != nil {
		// co the persist vao db transaction_revert || kafka
		if err == nil {
			err = err1
		}
		errx := errorx.SuspendedErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
	}
	return nil, nil
}
