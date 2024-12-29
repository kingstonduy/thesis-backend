package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
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
func (repo *transactionRepo) Insert(ctx context.Context, tr domain.TransactionEntity) error {
	logger.Info(ctx, "Insert transaction start")
	defer logger.Info(ctx, "Insert transaction end")

	sqlQuery := `
        INSERT INTO public."TRANSACTION" (
            "TRANSACTION_ID",
            "STATUS",
            "PROCESSING",
            "CREATED_AT",
            "UPDATED_AT"
        )
        VALUES (
            $1,
            $2,
            $3,
            CURRENT_TIMESTAMP,
            CURRENT_TIMESTAMP
        );
    `

	_, err := repo.db.DB.Exec(ctx, sqlQuery,
		tr.TransactionID,
		tr.Status,
		tr.Processing,
	)
	if err != nil {
		logger.Errorf(ctx, "Failed to insert transaction: %v", err)
		return err
	}

	return nil
}

// SelectByTransactionEntityID implements domain.ITransactionRepo.
func (repo *transactionRepo) SelectByTransactionEntityID(ctx context.Context, id string) (domain.TransactionEntity, error) {
	logger.Info(ctx, "Select transaction by ID start")
	defer logger.Info(ctx, "Select transaction by ID end")

	sqlQuery := `
        SELECT 
            "TRANSACTION_ID",
            "STATUS",
            "PROCESSING",
            "CREATED_AT",
            "UPDATED_AT"
        FROM 
            public."TRANSACTION"
        WHERE 
            "TRANSACTION_ID" = $1;
    `

	var tr domain.TransactionEntity
	err := repo.db.DB.QueryRow(ctx, sqlQuery, id).Scan(
		&tr.TransactionID,
		&tr.Status,
		&tr.Processing,
		&tr.CreatedAt,
		&tr.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Warnf(ctx, "Transaction not found for ID: %s", id)
			return tr, nil
		}
		logger.Errorf(ctx, "Failed to select transaction: %v", err)
		return tr, err
	}

	return tr, nil
}

// Update implements domain.ITransactionRepo.
func (repo *transactionRepo) Update(ctx context.Context, tr domain.TransactionEntity) error {
	logger.Info(ctx, "Update transaction start")
	defer logger.Info(ctx, "Update transaction end")

	sqlQuery := `
        UPDATE public."TRANSACTION"
        SET 
            "STATUS" = $1,
            "PROCESSING" = $2,
            "UPDATED_AT" = CURRENT_TIMESTAMP
        WHERE 
            "TRANSACTION_ID" = $3;
    `

	res, err := repo.db.DB.Exec(ctx, sqlQuery,
		tr.Status,
		tr.Processing,
		tr.TransactionID,
	)
	if err != nil {
		logger.Errorf(ctx, "Failed to update transaction: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Errorf(ctx, "Failed to fetch rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		err = fmt.Errorf("no rows affected when updating transaction")
		logger.Warnf(ctx, err.Error())
		return err
	}

	return nil
}
