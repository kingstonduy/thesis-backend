package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/comment-service/internal/domain"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
)

type handler struct {
	repo domain.ICommentRepo
}

func NewAddCommentHandler(
	repo domain.ICommentRepo,
) domain.IAddCommentHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.AddCommentHandler.
func (h *handler) Handle(ctx context.Context, req *domain.AddCommentRequest) (res *domain.AddCommentResponse, err error) {
	logger.Info(ctx, "AddCommentHandler start")
	defer logger.Info(ctx, "AddCommentHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC AddCommentHandler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	err = h.repo.AddComment(ctx, domain.AddCommentParams{
		ProductID:        req.ProductID,
		UserID:           req.UserID,
		ReviewEvaluation: req.Rating,
		ReviewDetail:     req.Content,
	})
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}
	return res, nil
}
