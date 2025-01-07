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

func NewGetCartHandler(
	repo domain.ICartRepo,
) domain.IGetCartHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IGetCartHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetCartRequest) (res *domain.GetCartResponse, err error) {
	logger.Info(ctx, "IGetCartHandler start")
	defer logger.Info(ctx, "IGetCartHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC IGetCartHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	params := domain.GetCartParamsIn{
		UserID: req.UserID,
	}
	params.UserID = req.UserID

	cartItems, err := h.repo.GetCart(ctx, params)
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	res = &domain.GetCartResponse{
		CartItems: cartItems.CartItems,
	}

	return res, nil
}
