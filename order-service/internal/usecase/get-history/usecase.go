package usecase

import (
	"context"

	"github.com/kingstonduy/order-service/internal/domain"
)

type handler struct {
}

func NewGetHistoryHandler() domain.GetHistoryHandler {
	return &handler{}
}

// Handle implements domain.GetHistoryHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetHistoryRequest) (res *domain.GetHistoryResponse, err error) {
	panic("unimplemented")
}
