package postgres

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
)

type cartRepoImpl struct {
	db *configuration.PostgresCon
}

func NewOrderRepo(db *configuration.PostgresCon) domain.IOrderRepo {
	return &cartRepoImpl{
		db: db,
	}
}

// GetHistory implements domain.IOrderRepo.
func (c *cartRepoImpl) GetHistory(ctx context.Context, params domain.GetHistoryParamIn) (domain.GetHistoryResponse, error) {
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
            oi."PAYMENT_STATUS"
        FROM 
            public."ORDER_ITEM" oi
        INNER JOIN 
            public."PRODUCT" p
        ON 
            oi."PRODUCT_ID" = p."PRODUCT_ID"
        WHERE 
            oi."USER_ID" = $1
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
func (c *cartRepoImpl) Insert(ctx context.Context, orderEntity domain.OrderEntity) error {
	logger.Info(ctx, "Insert start")
	defer logger.Info(ctx, "Insert end")

	sqlQuery := `
        INSERT INTO public."ORDER" (
            "PRODUCT_ID",
            "USER_ID",
            "TRANSACTION_ID",
            "DELIVERY_STATUS",
            "PAYMENT_STATUS",
            "CREATED_AT",
            "UPDATED_AT"
        )
        VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
            CURRENT_TIMESTAMP,
            CURRENT_TIMESTAMP
        );
    `

	// Insert the order record
	res, err := c.db.DB.Exec(ctx, sqlQuery,
		orderEntity.ProductID,
		orderEntity.UserID,
		orderEntity.TransactionID,
		orderEntity.DeliveryStatus,
		orderEntity.PaymentStatus,
	)
	if err != nil {
		logger.Errorf(ctx, "Failed to insert order: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Errorf(ctx, "Failed to fetch rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		err = fmt.Errorf("no rows affected when inserting order")
		logger.Errorf(ctx, err.Error())
		return err
	}

	return nil
}

// Update implements domain.IOrderRepo.
func (c *cartRepoImpl) Update(ctx context.Context, orderEntity domain.OrderEntity) error {
	logger.Info(ctx, "Update start")
	defer logger.Info(ctx, "Update end")

	sqlQuery := `
        UPDATE public."ORDER"
        SET 
            "PRODUCT_ID" = $1,
            "USER_ID" = $2,
            "TRANSACTION_ID" = $3,
            "DELIVERY_STATUS" = $4,
            "PAYMENT_STATUS" = $5,
            "UPDATED_AT" = CURRENT_TIMESTAMP
        WHERE 
            "ORDER_ID" = $6;
    `

	// Update the order record
	res, err := c.db.DB.Exec(ctx, sqlQuery,
		orderEntity.ProductID,      // $1
		orderEntity.UserID,         // $2
		orderEntity.TransactionID,  // $3
		orderEntity.DeliveryStatus, // $4
		orderEntity.PaymentStatus,  // $5
		orderEntity.OrderID,        // $6
	)
	if err != nil {
		logger.Errorf(ctx, "Failed to update order: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Errorf(ctx, "Failed to fetch rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		err = fmt.Errorf("no rows affected when updating order")
		logger.Errorf(ctx, err.Error())
		return err
	}

	return nil
}
