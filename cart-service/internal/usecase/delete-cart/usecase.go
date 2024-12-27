package usecase

import (
	"context"

	"github.com/kingstonduy/cart-service/internal/domain"
)

type handler struct {
}

func NewDeleteCartItemHandler() domain.DeleteCartItemHandler {
	return &handler{}
}

// Handle implements domain.DeleteCartItemHandler.
func (h *handler) Handle(ctx context.Context, req *domain.DeleteCartItemRequest) (res *domain.DeleteCartItemResponse, err error) {
	panic("unimplemented")
}
