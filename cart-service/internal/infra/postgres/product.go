package postgres

import (
	"context"
	"fmt"

	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
	gensql "github.com/kingstonduy/cart-service/internal/pkg/gen_sql"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
)

type productRepo struct {
	db *configuration.PostgresCon
}

func NewProductRepo(db *configuration.PostgresCon) domain.IProductRepo {
	return &productRepo{
		db: db,
	}
}

// Insert implements domain.IProductRepo.
func (repo *productRepo) Insert(ctx context.Context, entity domain.ProductCdc) error {
	sqlQuery, _ := gensql.GenInsertSql("PRODUCT", entity)

	logger.Infof(ctx, "sql query: %s", sqlQuery)

	resp, err := repo.db.DB.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	affectedRows, err := resp.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return fmt.Errorf(errorx.ErrorMessageNoRowAffected)
	}

	return nil
}

// Update implements domain.IProductRepo.
func (p *productRepo) Update(ctx context.Context, cols map[string]interface{}, condition map[string]interface{}) error {
	panic("unimplemented")
}
