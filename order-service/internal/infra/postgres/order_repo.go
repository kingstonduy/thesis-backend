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

type orderRepoImpl struct {
	db *configuration.PostgresCon
}

func NewOrderRepo(db *configuration.PostgresCon) domain.IOrderRepo {
	return &orderRepoImpl{
		db: db,
	}
}

// GetHistory implements domain.IOrderRepo.
func (c *orderRepoImpl) GetHistory(ctx context.Context, params domain.GetHistoryParamIn) (domain.GetHistoryResponse, error) {
	logger.Info(ctx, "GetHistory start")
	defer logger.Info(ctx, "GetHistory end")

	// SQL query to fetch order history for the user
	sqlQuery := `
    SELECT 
        oi."PRODUCT_ID",
        p."PRODUCT_IMAGE",
        p."PRODUCT_NAME",
        oi."ORDER_ID",
        oi."DELIVERY_STATUS",
        oi."PAYMENT_STATUS",
        t."STATUS" as "TRANSACTION_STATUS",
        oi."CREATED_AT",
        oi."UPDATED_AT"
    FROM 
        public."ORDER_ITEM" oi
    INNER JOIN 
        public."PRODUCT" p
    ON 
        oi."PRODUCT_ID" = p."PRODUCT_ID"
    INNER JOIN 
        public."TRANSACTION" t
    ON 
        oi."TRANSACTION_ID" = t."TRANSACTION_ID" -- Join with TRANSACTION table
    WHERE 
        oi."USER_ID" = $1
        AND oi."DELIVERY_STATUS" IS NOT NULL
    ORDER BY 
        oi."CREATED_AT" DESC;
    `

	// Prepare response
	var response domain.GetHistoryResponse
	rows, err := c.db.DB.Query(ctx, sqlQuery, params.UserID)
	if err != nil {
		logger.Errorf(ctx, "Failed to fetch order history: %v", err)
		return response, err
	}
	defer rows.Close()

	for rows.Next() {
		var detail domain.GetHistoryResponseDetail
		err := rows.Scan(
			&detail.ProductID,
			&detail.ProductImage,
			&detail.ProductName,
			&detail.OrderID,
			&detail.DeliveryStatus,
			&detail.PaymentStatus,
			&detail.TransactionStatus,
			&detail.CreatedAt,
			&detail.UpdatedAt,
		)
		if err != nil {
			logger.Errorf(ctx, "Failed to scan row: %v", err)
			return response, err
		}
		response.Details = append(response.Details, detail)
	}

	// Check for any errors during iteration
	if rows.Err() != nil {
		logger.Errorf(ctx, "Error iterating over rows: %v", rows.Err())
		return response, rows.Err()
	}

	return response, nil
}

// Insert implements domain.IOrderRepo.
func (repo *orderRepoImpl) Insert(ctx context.Context, entity domain.OrderEntity) error {
	logger.Info(ctx, "Insert ORDER_ITEM starts")
	defer logger.Info(ctx, "Insert ORDER_ITEM ends")

	sqlQuery, err := gensql.GenInsertSql("ORDER_ITEM", entity)
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

// Update implements domain.IOrderRepo.
func (repo *orderRepoImpl) Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error {
	logger.Info(ctx, "Update ORDER_ITEM starts")
	defer logger.Info(ctx, "Update ORDER_ITEM ends")

	sqlQuery, err := gensql.GenUpdateSql("ORDER_ITEM", cols, conditions)
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
