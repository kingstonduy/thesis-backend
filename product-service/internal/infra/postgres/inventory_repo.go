package postgres

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
	gensql "github.com/kingstonduy/product-service/internal/pkg/gen_sql"
)

type inventoryRepoImlp struct {
	db *configuration.PostgresCon
}

func NewInventoryRepoImpl(
	db *configuration.PostgresCon,
) domain.IInventoryRepo {
	return &inventoryRepoImlp{
		db: db,
	}
}

// SelectByID implements domain.IInventoryRepo.
func (repo *inventoryRepoImlp) SelectByProductID(ctx context.Context, productID string) (entity domain.Inventory, err error) {
	logger.Info(ctx, "Select INVENTORY ByProductID start")
	defer logger.Info(ctx, "Select INVENTORY ByProductID start")

	sqlQuery := `
        select * from "INVENTORY" where "PRODUCT_ID"=$1;
    `

	if err = repo.db.DB.Get(ctx, &entity, sqlQuery, productID); err != nil {
		return entity, err
	}
	return entity, nil
}

// Insert implements domain.IInventoryRepo.
func (repo *inventoryRepoImlp) Insert(ctx context.Context, entity domain.Inventory) error {
	logger.Info(ctx, "Insert INVENTORY starts")
	defer logger.Info(ctx, "Insert INVENTORY ends")

	sqlQuery, err := gensql.GenInsertSql("INVENTORY", entity)
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

// Update implements domain.IInventoryRepo.
func (repo *inventoryRepoImlp) Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error {
	logger.Info(ctx, "Update INVENTORY starts")
	defer logger.Info(ctx, "Update INVENTORY ends")

	sqlQuery, err := gensql.GenUpdateSql("INVENTORY", cols, conditions)
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
