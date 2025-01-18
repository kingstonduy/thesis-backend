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
) domain.IUpdateCartItemHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IUpdateCartItemHandler.
func (h *handler) Handle(ctx context.Context, req *domain.UpdateCartItemRequest) (res *domain.UpdateCartItemResponse, err error) {
	logger.Info(ctx, "IUpdateCartItemHandler start")
	defer logger.Info(ctx, "IUpdateCartItemHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC IUpdateCartItemHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	if req.CartItemQuantity == 0 { // delete cartITem
		err = h.repo.DeleteById(ctx, req.CartItemID)
		if err != nil {
			errx := errorx.OutboundErrorWithDetails(err.Error(), "")
			logger.Error(ctx, errx.Error())
			return nil, errx
		}
		return nil, nil
	} else {
		cols := map[string]interface{}{}
		conditions := map[string]interface{}{}

		cols["CART_ITEM_QUANTITY"] = req.CartItemQuantity
		conditions["CART_ITEM_ID"] = req.CartItemID

		err = h.repo.Update(ctx, cols, conditions)
		if err != nil {
			errx := errorx.OutboundErrorWithDetails(err.Error(), "")
			logger.Error(ctx, errx.Error())
			return nil, errx
		}

		return nil, nil
	}

}
