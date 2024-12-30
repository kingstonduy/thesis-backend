package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

const (
	DEFAULT_USER_IMAGE = "https://avatars.githubusercontent.com/u/77771338?v=4"
	DEFAULT_FIRST_NAME = "John"
	DEFAULT_LAST_NAME  = "Doe"
)

type handler struct {
	repo domain.IUserRepo
}

func NewRegisterHandler(
	repo domain.IUserRepo,
) domain.IRegisterHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IRegisterHandler.
func (h *handler) Handle(ctx context.Context, req *domain.RegisterRequest) (res *domain.RegisterResponse, err error) {
	logger.Info(ctx, "Register handler start")
	defer logger.Info(ctx, "Register handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Register handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	now := time.Now()

	entity := domain.UserEntity{
		UserID:       uuid.New().String(),
		UserName:     req.UserName,
		UserPassword: HashPassword(req.Password),
		UserImage:    DEFAULT_USER_IMAGE,
		FirstName:    DEFAULT_FIRST_NAME,
		LastName:     DEFAULT_LAST_NAME,
		DateOfBirth:  req.DateOfBirth,
		Gender:       req.Gender,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		Street:       req.Street,
		City:         req.City,
		CityCode:     req.CityCode,
		District:     req.District,
		DistrictCode: req.DistrictCode,
		Ward:         req.Ward,
		WardCode:     req.WardCode,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if err = h.repo.Insert(ctx, entity); err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	return res, nil
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) string {
	// Generate a salted hash for the password
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
