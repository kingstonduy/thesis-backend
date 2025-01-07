package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
)

type handler struct {
	repo domain.ICartRepo
}

func NewIDeleteCartItemHandler(
	repo domain.ICartRepo,
) domain.IDeleteCartItemHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IDeleteCartItemHandler.
func (h *handler) Handle(ctx context.Context, req *domain.DeleteCartItemsRequest) (res *domain.DeleteCartItemsResponse, err error) {
	logger.Info(ctx, "IDeleteCartItemHandler start")
	defer logger.Info(ctx, "IDeleteCartItemHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC IDeleteCartItemHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	var paramsIn domain.DeleteCartItemParamsIn
	for _, detail := range req.Details {

		paramsIn.Details = append(paramsIn.Details, domain.DeleteCartItemParamsInDetails{
			CartItemID: detail.CartItemID,
		})
	}
	err = h.repo.DeleteCartItemsByID(ctx, paramsIn)
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}
	return nil, nil
}
