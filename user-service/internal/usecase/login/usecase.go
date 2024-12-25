package usecase

import (
	"context"

	"github.com/kingstonduy/user-service/internal/domain"
)

type handler struct {
	repo domain.IUserRepo
}

func NewLoginHandler(
	repo domain.IUserRepo,
) domain.IRLoginHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IRLoginHandler.
func (h *handler) Handle(ctx context.Context, req *domain.LoginRequest) (res *domain.LoginResponse, err error) {
	panic("unimplemented")
}
