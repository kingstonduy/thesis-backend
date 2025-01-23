package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
)

const (
	DEFAULT_USER_IMAGE = "https://avatars.githubusercontent.com/u/77771338?v=4"
)

type handler struct {
	locationRepo domain.ILocationRepo
	repo         domain.IUserRepo
}

func NewRegisterHandler(
	locationRepo domain.ILocationRepo,
	repo domain.IUserRepo,
) domain.IRegisterHandler {
	return &handler{
		locationRepo: locationRepo,
		repo:         repo,
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

	var cityName string
	var districtName string
	var wardName string

	g := new(errgroup.Group)
	g.Go(func() error {
		cityName, err = h.locationRepo.GetCity(ctx, req.CityCode)
		return err
	})
	g.Go(func() error {
		districtName, err = h.locationRepo.GetDistrict(ctx, req.DistrictCode)
		return err
	})
	g.Go(func() error {
		wardName, err = h.locationRepo.GetWard(ctx, req.WardCode)
		return err
	})

	// wait for the subscription result, return error if present
	if err := g.Wait(); err != nil {
		return nil, errorx.FailedWithDetails(err.Error(), "")
	}

	entity := domain.UserEntity{
		UserID:       uuid.New().String(),
		UserName:     req.UserName,
		UserPassword: HashPassword(req.Password),
		UserImage:    DEFAULT_USER_IMAGE,
		FirstName:    gofakeit.FirstName(),
		LastName:     gofakeit.LastName(),
		DateOfBirth:  req.DateOfBirth,
		Gender:       req.Gender,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		Street:       req.Street,
		City:         cityName,
		CityCode:     req.CityCode,
		District:     districtName,
		DistrictCode: req.DistrictCode,
		Ward:         wardName,
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
