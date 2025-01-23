package postgres

import (
	"context"

	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/user-service/internal/bootstrap"
	"github.com/kingstonduy/user-service/internal/domain"
)

type locationRepo struct {
	db *configuration.PostgresCon
}

func NewLocationRepo(db *configuration.PostgresCon) domain.ILocationRepo {
	return &locationRepo{
		db: db,
	}
}

// GetCity implements domain.ILocationRepo.
func (repo *locationRepo) GetCity(ctx context.Context, cityCode string) (s string, err error) {
	logger.Info(ctx, "GetCity name start")
	defer logger.Info(ctx, "GetCity name end")

	var entity domain.CityEntity
	sqlQuery := `
        select * from "CITY" where "CITY_CODE"=$1;
    `
	if err = repo.db.DB.Get(ctx, &entity, sqlQuery, entity); err != nil {
		return "", err
	}

	return entity.Name, nil
}

// GetDistrict implements domain.ILocationRepo.
func (repo *locationRepo) GetDistrict(ctx context.Context, districtCode string) (s string, err error) {
	logger.Info(ctx, "Get DISTRICT name start")
	defer logger.Info(ctx, "Get DISTRICT name end")

	var entity domain.DistrictEntity
	sqlQuery := `
        select * from "DISTRICT" where "DISTRICT_CODE"=$1;
    `
	if err = repo.db.DB.Get(ctx, &entity, sqlQuery, entity); err != nil {
		return "", err
	}

	return entity.Name, nil
}

// GetWard implements domain.ILocationRepo.
func (repo *locationRepo) GetWard(ctx context.Context, wardCode string) (s string, err error) {
	logger.Info(ctx, "Get WARD name start")
	defer logger.Info(ctx, "Get WARD name end")

	var entity domain.WardEntity
	sqlQuery := `
        select * from "WARD" where "WARD_CODE"=$1;
    `
	if err = repo.db.DB.Get(ctx, &entity, sqlQuery, entity); err != nil {
		return "", err
	}

	return entity.Name, nil
}
