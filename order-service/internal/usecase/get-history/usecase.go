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

func NewGetHistoryHandler(
	orderRepo domain.IOrderRepo,
	transactionRepo domain.ITransactionRepo,
) domain.IGetHistoryHandler {
	return &handler{
		orderRepo:       orderRepo,
		transactionRepo: transactionRepo,
	}
}

// Handle implements domain.GetHistoryHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetHistoryRequest) (res *domain.GetHistoryResponse, err error) {
	logger.Info(ctx, "GetHistoryHandler start")
	defer logger.Info(ctx, "GetHistoryHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC GetHistoryHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	resp, err := h.orderRepo.GetHistory(ctx, domain.GetHistoryParamIn{UserID: req.UserID})
	if err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	res = &resp

	return res, nil
}
