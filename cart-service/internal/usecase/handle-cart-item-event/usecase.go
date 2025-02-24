package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/logger"
)

type handler struct {
	cartItemRepo domain.IReadCartItemRepo
}

func NewCartItemEventHandler(
	cartItemRepo domain.IReadCartItemRepo,
) domain.ICartItemEventHandler {
	return &handler{
		cartItemRepo: cartItemRepo,
	}
}

// Handle implements domain.CartItemEventHandler.
func (h *handler) Handle(ctx context.Context, req domain.Event[*domain.CartItemEvent]) (res *domain.CartItemEventRes, err error) {
	logger.Infof(ctx, "CartItemEventHandler start")
	defer logger.Infof(ctx, "CartItemEventHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC CartItemEventHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	if req.Payload.Deleted {
		err = h.cartItemRepo.Delete(ctx, req.Payload.CartItemID)
		if err != nil {
			logger.Error(ctx, err.Error())
			return nil, err
		}
	} else {
		entity := domain.ReadCartItemEntity{
			CartItemID:       req.Payload.CartItemID,
			UserID:           req.Payload.UserID,
			ProductID:        req.Payload.ProductID,
			CartItemQuantity: req.Payload.CartItemQuantity,
		}
		err = h.cartItemRepo.Upsert(ctx, entity)
		if err != nil {
			logger.Error(ctx, err.Error())
			return nil, err
		}
	}
	return nil, nil
}
