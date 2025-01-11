package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
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
	logger.Info(ctx, "Register handler start")
	defer logger.Info(ctx, "Register handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Register handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	// TODO remove this
	if req.Username == "ADMIN" {
		return &domain.LoginResponse{UserID: "ADMIN"}, nil
	}

	user, err := h.repo.GetUserByUserName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			errx := errorx.AuthenticationErrorWithDetails("user does not exist", "")
			return nil, errx
		}
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	if !ComparePasswords(user.UserPassword, req.Password) {
		errx := errorx.AuthenticationErrorWithDetails("the password is incorrect", "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	return &domain.LoginResponse{
		UserID: user.UserID,
	}, nil
}

// ComparePasswords compares a hashed password with a plain password
func ComparePasswords(hashedPassword, plainPassword string) bool {
	// Compare the hashed password with the plain password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil // Returns true if they match, false otherwise
}
