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

func NewDeleteUserCartHandler(
	repo domain.ICartRepo,
) domain.DeleteUserCartHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.DeleteUserCartHandler.
func (h *handler) Handle(ctx context.Context, req *domain.DeleteUserCartRequest) (res *domain.DeleteUserCartResponse, err error) {
	logger.Info(ctx, "DeleteUserCartHandler start")
	defer logger.Info(ctx, "DeleteUserCartHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC DeleteUserCartHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	err = h.repo.DeleteUserCart(ctx, domain.DeleteUserCartParams{UserID: req.UserID})
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}
	return nil, nil
}
