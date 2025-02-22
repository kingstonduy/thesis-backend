package postgres

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/cache"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
	gensql "github.com/kingstonduy/product-service/internal/pkg/gen_sql"
)

const (
	REDIS_KEY_GET_PRODUCTS_GENDER   = "PRODUCTS_GENDER"
	REDIS_KEY_GET_PRODUCTS_CATEGORY = "PRODUCTS_CATEGORY"
	REDIS_KEY_GET_PRODUCTS          = "PRODUCTS_ALL"
	REDIS_KEY_GET_PRODUCT_ID        = "PRODUCT_ID"
)

type productRepoImlp struct {
	db          *configuration.PostgresCon
	redisCLient cache.CacheClient
}

func NewProductRepoImpl(
	db *configuration.PostgresCon,
	redisCLient cache.CacheClient,
) domain.IWriteProductRepo {
	return &productRepoImlp{
		db:          db,
		redisCLient: redisCLient,
	}
}

// Insert implements domain.IProductRepo.
func (repo *productRepoImlp) Insert(ctx context.Context, entity interface{}) error {
	logger.Info(ctx, "Insert PRODUCT starts")
	defer logger.Info(ctx, "Insert PRODUCT ends")

	sqlQuery, err := gensql.GenInsertSql("PRODUCT", entity)
	if err != nil {
		return err
	}

	logger.Info(ctx, sqlQuery)

	res, err := repo.db.DB.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return fmt.Errorf(errorx.ErrorMessageNoRowAffected)
	}

	return nil
}

// Update implements domain.IProductRepo.
func (repo *productRepoImlp) Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error {
	logger.Info(ctx, "Update PRODUCT starts")
	defer logger.Info(ctx, "Update PRODUCT ends")

	sqlQuery, err := gensql.GenUpdateSql("PRODUCT", cols, conditions)
	if err != nil {
		return err
	}

	logger.Info(ctx, sqlQuery)

	res, err := repo.db.DB.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return fmt.Errorf(errorx.ErrorMessageNoRowAffected)
	}

	return nil
}
