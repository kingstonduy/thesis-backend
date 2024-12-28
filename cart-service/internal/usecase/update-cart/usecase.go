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

func NewUpdateCartHandler(
	repo domain.ICartRepo,
) domain.UpdateCartItemHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.UpdateCartItemHandler.
func (h *handler) Handle(ctx context.Context, req *domain.UpdateCartItemRequest) (res *domain.UpdateCartItemResponse, err error) {
	logger.Info(ctx, "UpdateCartItemHandler start")
	defer logger.Info(ctx, "UpdateCartItemHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC UpdateCartItemHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	if req.CartItemQuantity == 0 { // delete cartITem
		err = h.repo.DeleteCartItem(ctx, domain.DeleteCartItemParams{CartItemID: req.CartItemID})
		if err != nil {
			errx := errorx.OutboundErrorWithDetails(err.Error(), "")
			logger.Error(ctx, errx.Error())
			return nil, errx
		}
		return nil, nil

	} else {
		err = h.repo.UpdateCartItem(ctx, domain.UpdateCartItemParams{
			CartItemID:       req.CartItemID,
			CartItemQuantity: req.CartItemQuantity,
		})

		if err != nil {
			errx := errorx.OutboundErrorWithDetails(err.Error(), "")
			logger.Error(ctx, errx.Error())
			return nil, errx
		}

		return res, nil
	}

}
