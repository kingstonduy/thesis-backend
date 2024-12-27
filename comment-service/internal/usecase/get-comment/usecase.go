package usecase

import (
	"context"

	"github.com/kingstonduy/comment-service/internal/domain"
)

type handler struct {
}

func NewGetCommentHandler() domain.GetCommentHandler {
	return &handler{}
}

// Handle implements domain.GetCommentHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetCommentRequest) (res *domain.GetCommentResponse, err error) {
	panic("unimplemented")
}
