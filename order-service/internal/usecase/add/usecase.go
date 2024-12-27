package usecase

import (
	"context"

	"github.com/kingstonduy/order-service/internal/domain"
)

type handler struct {
}

func NewAddHandler() domain.AddHandler {
	return &handler{}
}

// Handle implements domain.AddHandler.
func (h *handler) Handle(ctx context.Context, req *domain.AddRequest) (res *domain.AddResponse, err error) {
	panic("unimplemented")
}
