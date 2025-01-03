package postgres

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/cache"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
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
) domain.IProductRepo {
	return &productRepoImlp{
		db:          db,
		redisCLient: redisCLient,
	}
}

// GetProductByCategory implements domain.IProductRepo.
func (repo *productRepoImlp) GetProductByCategory(ctx context.Context, category string) (products []domain.ProductEntity, err error) {
	logger.Info(ctx, "GetProductByGender start")
	defer logger.Info(ctx, "GetProductByGender end")

	key := REDIS_KEY_GET_PRODUCTS_CATEGORY + "-" + category
	dur, err := repo.redisCLient.Get(ctx, key, &products)
	if dur == -2 || err != nil {
		logger.Errorf(ctx, "cache miss key=%s", key)
		sqlQuery := `
        SELECT 
        "PRODUCT_ID", "PRODUCT_NAME", "PRODUCT_DESCRIPTION", "PRODUCT_IMAGE", 
        "PRODUCT_QUANTITY", "PRODUCT_PRICE", "CREATED_AT", "UPDATED_AT", 
        "PRODUCT_CATEGORY", "GENDER", "AVERAGE_RATING", "TOTAL_RATING"
        FROM public."PRODUCT"
        where "PRODUCT_CATEGORY"=$1;
    `

		rows, err := repo.db.DB.Query(ctx, sqlQuery, category)
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
				&product.TotalRating,
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

		go func() {
			err = repo.redisCLient.Set(context.Background(), key, products, 0)
			if err != nil {
				logger.Error(ctx, "Failed to set cache key=%s, error: %v", key, err)
			}
		}()
		return products, nil
	}

	return products, nil
}

// GetProductByGender implements domain.IProductRepo.
func (repo *productRepoImlp) GetProductByGender(ctx context.Context, gender string) (products []domain.ProductEntity, err error) {
	logger.Info(ctx, "GetProductByGender start")
	defer logger.Info(ctx, "GetProductByGender end")

	key := REDIS_KEY_GET_PRODUCTS_GENDER + "-" + gender
	dur, err := repo.redisCLient.Get(ctx, key, &products)
	if dur == -2 || err != nil {
		logger.Errorf(ctx, "cache miss key=%s", key)
		sqlQuery := `
        SELECT 
        "PRODUCT_ID", "PRODUCT_NAME", "PRODUCT_DESCRIPTION", "PRODUCT_IMAGE", 
        "PRODUCT_QUANTITY", "PRODUCT_PRICE", "CREATED_AT", "UPDATED_AT", 
        "PRODUCT_CATEGORY", "GENDER", "AVERAGE_RATING", "TOTAL_RATING"
        FROM public."PRODUCT"
        where "GENDER"=$1;
    `

		rows, err := repo.db.DB.Query(ctx, sqlQuery, gender)
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
				&product.TotalRating,
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

		go func() {
			err = repo.redisCLient.Set(context.Background(), key, products, 0)
			if err != nil {
				logger.Error(ctx, "Failed to set cache key=%s, error: %v", key, err)
			}
		}()
		return products, nil
	}

	return products, nil
}

// GetAllProduct implements domain.IProductRepo.
func (repo *productRepoImlp) GetAllProduct(ctx context.Context) (products []domain.ProductEntity, err error) {
	logger.Info(ctx, "GetAllProduct start")
	defer logger.Info(ctx, "GetAllProduct end")

	dur, err := repo.redisCLient.Get(ctx, REDIS_KEY_GET_PRODUCTS, &products)
	if dur == -2 || err != nil {
		logger.Errorf(ctx, "cache miss key=%s", REDIS_KEY_GET_PRODUCTS)
		sqlQuery := `
        SELECT 
        "PRODUCT_ID", "PRODUCT_NAME", "PRODUCT_DESCRIPTION", "PRODUCT_IMAGE", 
        "PRODUCT_QUANTITY", "PRODUCT_PRICE", "CREATED_AT", "UPDATED_AT", 
        "PRODUCT_CATEGORY", "GENDER", "AVERAGE_RATING","TOTAL_RATING"
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
				&product.TotalRating,
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

		go func() {
			err = repo.redisCLient.Set(context.Background(), REDIS_KEY_GET_PRODUCTS, products, 0)
			if err != nil {
				logger.Error(ctx, "Failed to set cache key=%s, error: %v", REDIS_KEY_GET_PRODUCTS, err)
			}
		}()

		return products, nil
	}

	return products, nil
}

// GetProductByID implements domain.IProductRepo.
func (repo *productRepoImlp) GetProductByID(ctx context.Context, productID string) (entity domain.ProductEntity, err error) {
	logger.Info(ctx, "GetProductByID start")
	defer logger.Info(ctx, "GetProductByID end")

	key := REDIS_KEY_GET_PRODUCT_ID + "-" + productID

	dur, err := repo.redisCLient.Get(ctx, key, &entity)
	if dur == -2 || err != nil {
		logger.Errorf(ctx, "cache miss key=%s", key)
		sqlQuery := `
        select * from "PRODUCT" where "PRODUCT_ID"=$1;
    `

		if err = repo.db.DB.Get(ctx, &entity, sqlQuery, productID); err != nil {
			return entity, err
		}

		go func() {
			err = repo.redisCLient.Set(context.Background(), key, entity, 0)
			if err != nil {
				logger.Error(ctx, "Failed to set cache key=%s, error: %v", REDIS_KEY_GET_PRODUCTS, err)
			}
		}()
		return entity, nil
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
            "GENDER"=$8,
            "TOTAL_RATING"=$10,
        WHERE 
            "PRODUCT_ID"=$1 AND
            "UPDATED_AT"=$9
    `
	res, err := repo.db.DB.Exec(ctx, sqlQuery,
		product.ProductID, product.ProductName, product.ProductDescription,
		product.ProductImage, product.ProductQuantity, product.ProductPrice,
		product.ProductCategory, product.Gender, product.UpdatedAt, product.TotalRating,
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
