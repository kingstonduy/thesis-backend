package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/user-service/internal/domain"
	sql_util "github.com/kingstonduy/user-service/internal/pkg/utils/sql"
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
		UserID:       sql_util.SetString(uuid.New().String()),
		UserName:     sql_util.SetString(req.UserName),
		UserPassword: sql_util.SetString(req.Password),
		UserImage:    sql_util.SetString(DEFAULT_USER_IMAGE),
		FirstName:    sql_util.SetString(DEFAULT_FIRST_NAME),
		LastName:     sql_util.SetString(DEFAULT_LAST_NAME),
		DateOfBirth:  sql_util.SetString(req.DateOfBirth),
		Gender:       sql_util.SetString(req.Gender),
		Email:        sql_util.SetString(req.Email),
		PhoneNumber:  sql_util.SetString(req.PhoneNumber),
		Street:       sql_util.SetString(req.Street),
		City:         sql_util.SetString(req.City),
		CityCode:     sql_util.SetString(req.CityCode),
		District:     sql_util.SetString(req.District),
		DistrictCode: sql_util.SetString(req.DistrictCode),
		Ward:         sql_util.SetString(req.Ward),
		WardCode:     sql_util.SetString(req.WardCode),
		CreatedAt:    sql_util.SetTime(now),
		UpdatedAt:    sql_util.SetTime(now),
	}
	if err = h.repo.Insert(ctx, entity); err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	return res, nil
}
