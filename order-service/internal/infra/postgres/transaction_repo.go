package postgres

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
	gensql "github.com/kingstonduy/order-service/internal/pkg/gen_sql"
)

type transactionRepo struct {
	db *configuration.PostgresCon
}

func NewTransactionRepo(
	db *configuration.PostgresCon,
) domain.ITransactionRepo {
	return &transactionRepo{
		db: db,
	}
}

// Insert implements domain.ITransactionRepo.
func (repo *transactionRepo) Insert(ctx context.Context, entity domain.TransactionEntity) error {
	logger.Info(ctx, "Insert TRANSACTION starts")
	defer logger.Info(ctx, "Insert TRANSACTION ends")

	sqlQuery, err := gensql.GenInsertSql("TRANSACTION", entity)
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

// Update implements domain.ITransactionRepo.
func (repo *transactionRepo) Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error {
	logger.Info(ctx, "Update TRANSACTION starts")
	defer logger.Info(ctx, "Update TRANSACTION ends")

	sqlQuery, err := gensql.GenUpdateSql("TRANSACTION", cols, conditions)
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

// SelectByID implements domain.ITransactionRepo.
func (repo *transactionRepo) SelectByTransactionID(ctx context.Context, transactionID string) (entity domain.TransactionEntity, err error) {
	logger.Info(ctx, "Select TRANSACTION ByTransactionID start")
	defer logger.Info(ctx, "Select TRANSACTION ByTransactionID end")
	sqlQuery := `
            select * from "TRANSACTION" where "TRANSACTION_ID"=$1;
        `

	if err = repo.db.DB.Get(ctx, &entity, sqlQuery, transactionID); err != nil {
		return entity, err
	}

	return entity, nil
}
