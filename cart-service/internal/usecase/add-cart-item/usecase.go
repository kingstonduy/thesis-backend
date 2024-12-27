package usecase

import (
	"context"

	"github.com/kingstonduy/cart-service/internal/domain"
)

type handler struct {
}

func NewAddCartItemHandler() domain.AddCartItemHandler {
	return &handler{}
}

// Handle implements domain.AddCartItemHandler.
func (h *handler) Handle(ctx context.Context, req *domain.AddCartItemRequest) (res *domain.AddCartItemResponse, err error) {
	panic("unimplemented")
}
