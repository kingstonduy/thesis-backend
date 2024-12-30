package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/user-service/internal/domain"
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
		UserID:       req.UserID,
		UserName:     req.UserName,
		Gender:       req.Gender,
		Email:        req.Email,
		Street:       req.Street,
		City:         req.City,
		CityCode:     req.CityCode,
		District:     req.District,
		DistrictCode: req.DistrictCode,
		Ward:         req.Ward,
		WardCode:     req.WardCode,
		UpdatedAt:    now,
	}

	if err = h.repo.Update(ctx, entity); err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	return res, nil
}
