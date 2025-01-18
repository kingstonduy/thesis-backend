package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
)

type handler struct {
	repo domain.ICartRepo
}

func NewAddCartItemHandler(
	repo domain.ICartRepo,
) domain.IAddCartItemHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IAddCartItemHandler.
func (h *handler) Handle(ctx context.Context, req *domain.AddCartItemRequest) (res *domain.AddCartItemResponse, err error) {
	logger.Info(ctx, "IAddCartItemHandler start")
	defer logger.Info(ctx, "IAddCartItemHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC IAddCartItemHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	now := time.Now()

	cartItemEntity := domain.CartItem{
		CartItemID:       uuid.New().String(),
		UserID:           req.UserID,
		ProductID:        req.ProductID,
		CartItemQuantity: req.CartItemQuantity,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	if err := h.repo.Insert(ctx, cartItemEntity); err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	return nil, nil
}
