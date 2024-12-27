package usecase

import (
	"context"

	"github.com/kingstonduy/cart-service/internal/domain"
)

type handler struct {
}

func NewGetCartHandler() domain.GetCartHandler {
	return &handler{}
}

// Handle implements domain.GetCartHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetCartRequest) (res *domain.GetCartResponse, err error) {
	panic("unimplemented")
}
