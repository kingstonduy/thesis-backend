package usecase

import (
	"context"

	"github.com/kingstonduy/order-service/internal/domain"
)

type handler struct {
}

func NewGetCheckoutItemHandler() domain.GetCheckoutItemHandler {
	return &handler{}
}

// Handle implements domain.GetCheckoutItemHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetCheckoutItemRequest) (res *domain.GetCheckoutItemResponse, err error) {
	panic("unimplemented")
}
