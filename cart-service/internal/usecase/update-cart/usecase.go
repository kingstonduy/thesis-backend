package usecase

import (
	"context"

	"github.com/kingstonduy/cart-service/internal/domain"
)

type handler struct {
}

func NewUpdateCartHandler() domain.UpdateCartItemHandler {
	return &handler{}
}

// Handle implements domain.UpdateCartItemHandler.
func (h *handler) Handle(ctx context.Context, req *domain.UpdateCartItemRequest) (res *domain.UpdateCartItemResponse, err error) {
	panic("unimplemented")
}
