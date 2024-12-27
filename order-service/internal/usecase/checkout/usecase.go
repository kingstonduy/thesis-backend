package usecase

import (
	"context"

	"github.com/kingstonduy/order-service/internal/domain"
)

type handler struct {
}

func NewCheckoutHandler() domain.CheckoutHandler {
	return &handler{}
}

// Handle implements domain.CheckoutHandler.
func (h *handler) Handle(ctx context.Context, req *domain.CheckoutRequest) (res *domain.CheckoutResponse, err error) {
	panic("unimplemented")
}
