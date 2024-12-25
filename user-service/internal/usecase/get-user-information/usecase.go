package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/user-service/internal/domain"
)

type handler struct {
	repo domain.IUserRepo
}

func NewGetUserInformationhandler(
	repo domain.IUserRepo,
) domain.IGetUserInformationHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IGetUserInformationHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetUserInformationRequest) (res *domain.GetUserInformationResponse, err error) {
	logger.Info(ctx, "GetUserInformation handler start")
	defer logger.Info(ctx, "GetUserInformation handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC GetUserInformation handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	user, err := h.repo.GetUserByUserID(ctx, req.UserID)
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	res = &domain.GetUserInformationResponse{
		UserID:       user.UserID.String,
		UserName:     user.UserName.String,
		Email:        user.Email.String,
		PhoneNumber:  user.PhoneNumber.String,
		Gender:       user.Gender.String,
		DateOfBirth:  user.DateOfBirth.String,
		Street:       user.Street.String,
		City:         user.City.String,
		CityCode:     user.CityCode.String,
		District:     user.District.String,
		DistrictCode: user.DistrictCode.String,
		Ward:         user.Ward.String,
		WardCode:     user.WardCode.String,
	}
	return
}
