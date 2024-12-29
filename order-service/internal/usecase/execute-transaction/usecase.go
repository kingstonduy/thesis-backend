package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
	"github.com/kingstonduy/order-service/internal/pkg/utils"
	"github.com/pkg/errors"
)

type handler struct {
	orderRepo       domain.IOrderRepo
	transactionRepo domain.ITransactionRepo
	db              *configuration.PostgresCon
	cartOutbound    domain.ICartOutbound
	productOutbound domain.IProductOutbound
}

func NewExecuteTransactionHandler(
	orderRepo domain.IOrderRepo,
	transactionRepo domain.ITransactionRepo,
	db *configuration.PostgresCon,
	cartOutbound domain.ICartOutbound,
	productOutbound domain.IProductOutbound,
) domain.IExecuteTransactionHandler {
	return &handler{
		orderRepo:       orderRepo,
		transactionRepo: transactionRepo,
		db:              db,
		cartOutbound:    cartOutbound,
		productOutbound: productOutbound,
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
	err = h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
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

	// gọi cart
	_, err = h.cartOutbound.ExecuteTransaction(ctx, transport.Request[domain.CartExecuteTransactionRequest]{
		Data:  getCartExecuteTransactionRequest(*req),
		Trace: utils.GenRequestTrace(trace, "cart-service", ""),
	})
	if err != nil {
		transaction.Status = domain.FAILED_STATUS
		if err2 := h.transactionRepo.Update(ctx, transaction); err2 != nil {
			errx := errorx.FailedWithDetails(errors.Wrap(err, err2.Error()).Error(), "")
			logger.Error(ctx, errx.Error())
			return nil, errx
		}
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	// gọi product
	_, err = h.productOutbound.ExecuteTransaction(ctx, transport.Request[domain.ProductExecuteTransactionRequest]{
		Data:  getProductExecuteTransactionRequest(*req),
		Trace: utils.GenRequestTrace(trace, "product-service", ""),
	})
	if err != nil {
		transaction.Status = domain.FAILED_STATUS
		if err2 := h.transactionRepo.Update(ctx, transaction); err2 != nil {
			errx := errorx.FailedWithDetails(errors.Wrap(err, err2.Error()).Error(), "")
			logger.Error(ctx, errx.Error())
			return nil, errx
		}
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	// update transaction
	transaction.Status = domain.COMPLETE_STATUS
	if err = h.transactionRepo.Update(ctx, transaction); err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	return res, nil
}
