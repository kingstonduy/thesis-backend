package postgres

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
)

type productRepoImlp struct {
	db *configuration.PostgresCon
}

func NewProductRepoImpl(db *configuration.PostgresCon) domain.IProductRepo {
	return &productRepoImlp{
		db: db,
	}
}

// GetAllProduct implements domain.IProductRepo.
func (repo *productRepoImlp) GetAllProduct(ctx context.Context) (products []domain.ProductEntity, err error) {
	logger.Info(ctx, "GetAllProduct start")
	defer logger.Info(ctx, "GetAllProduct end")

	sqlQuery := `
        SELECT 
        "PRODUCT_ID", "PRODUCT_NAME", "PRODUCT_DESCRIPTION", "PRODUCT_IMAGE", 
        "PRODUCT_QUANTITY", "PRODUCT_PRICE", "CREATED_AT", "UPDATED_AT", 
        "PRODUCT_CATEGORY", "GENDER", "AVERAGE_RATING"
        FROM public."PRODUCT";
    `
	rows, err := repo.db.DB.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.ProductEntity
		err = rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.ProductDescription,
			&product.ProductImage,
			&product.ProductQuantity,
			&product.ProductPrice,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.ProductCategory,
			&product.Gender,
			&product.AvgRating,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if rows.Err() != nil {
		logger.Errorf(ctx, rows.Err().Error())
		return nil, rows.Err()
	}

	return products, nil
}

// GetProductByID implements domain.IProductRepo.
func (repo *productRepoImlp) GetProductByID(ctx context.Context, productID string) (entity domain.ProductEntity, err error) {
	logger.Info(ctx, "GetProductByID start")
	defer logger.Info(ctx, "GetProductByID end")

	sqlQuery := `
        select * from "PRODUCT" where "PRODUCT_ID"=$1;
    `

	if err = repo.db.DB.Get(ctx, &entity, sqlQuery, productID); err != nil {
		return entity, err
	}

	return entity, nil
}

// UpdateProductByID implements domain.IProductRepo.
func (repo *productRepoImlp) UpdateProductByID(ctx context.Context, product domain.ProductEntity) error {
	logger.Info(ctx, "GetProductByID start")
	defer logger.Info(ctx, "GetProductByID end")

	sqlQuery := `
        UPDATE public."PRODUCT"
        SET 
            "PRODUCT_NAME"=$2, 
            "PRODUCT_DESCRIPTION"=$3, 
            "PRODUCT_IMAGE"=$4, 
            "PRODUCT_QUANTITY"=$5, 
            "PRODUCT_PRICE"=$6, 
            "CREATED_AT"=CURRENT_TIMESTAMP, 
            "UPDATED_AT"=CURRENT_TIMESTAMP, 
            "PRODUCT_CATEGORY"=$7, 
            "GENDER"=$8
        WHERE 
            "PRODUCT_ID"=$1 AND
            "UPDATED_AT"=$9
    `
	res, err := repo.db.DB.Exec(ctx, sqlQuery,
		product.ProductID, product.ProductName, product.ProductDescription,
		product.ProductImage, product.ProductQuantity, product.ProductPrice,
		product.ProductCategory, product.Gender, product.UpdatedAt,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf(errorx.ErrorMessageNoRowAffected)
	}

	return nil
}
