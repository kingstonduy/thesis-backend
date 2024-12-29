package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/order-service/internal/domain"
)

type handler struct {
	orderRepo       domain.IOrderRepo
	transactionRepo domain.ITransactionRepo
}

func NewGetCheckoutItemHandler(
	orderRepo domain.IOrderRepo,
	transactionRepo domain.ITransactionRepo,
) domain.IGetCheckoutItemHandler {
	return &handler{
		orderRepo:       orderRepo,
		transactionRepo: transactionRepo,
	}
}

// Handle implements domain.GetCheckoutItemHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetCheckoutItemRequest) (res *domain.GetCheckoutItemResponse, err error) {
	logger.Info(ctx, "GetCheckoutItem start")
	defer logger.Info(ctx, "GetCheckoutItem end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC GetCheckoutItem %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	resp, err := h.orderRepo.GetCheckoutItem(ctx, domain.GetCheckoutItemParamIn{
		UserID: req.UserID,
	})
	if err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	res = &resp

	return res, nil
}
