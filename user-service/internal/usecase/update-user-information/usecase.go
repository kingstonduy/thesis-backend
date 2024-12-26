package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/user-service/internal/domain"
	sql_util "github.com/kingstonduy/user-service/internal/pkg/utils/sql"
)

type handler struct {
	repo domain.IUserRepo
}

func NewUpdateHandler(
	repo domain.IUserRepo,
) domain.IUpdateUserInformationHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IUpdateUserInformationHandler.
func (h *handler) Handle(ctx context.Context, req *domain.UpdateUserInformationRequest) (res *domain.UpdateUserInformationResponse, err error) {
	logger.Info(ctx, "Update User Information handler start")
	defer logger.Info(ctx, "Update User Information handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Update User Information handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	now := time.Now()

	entity := domain.UserEntity{
		UserID:       sql_util.SetString(req.UserID),
		UserName:     sql_util.SetString(req.UserName),
		Gender:       sql_util.SetString(req.Gender),
		Email:        sql_util.SetString(req.Email),
		Street:       sql_util.SetString(req.Street),
		City:         sql_util.SetString(req.City),
		CityCode:     sql_util.SetString(req.CityCode),
		District:     sql_util.SetString(req.District),
		DistrictCode: sql_util.SetString(req.DistrictCode),
		Ward:         sql_util.SetString(req.Ward),
		WardCode:     sql_util.SetString(req.WardCode),
		UpdatedAt:    sql_util.SetTime(now),
	}

	if err = h.repo.Update(ctx, entity); err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	return res, nil
}
