package postgres

import (
	"context"
	"fmt"
	"strings"

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
) domain.IProductRepo {
	return &productRepoImlp{
		db:          db,
		redisCLient: redisCLient,
	}
}

// GetAllProduct implements domain.IProductRepo.
func (repo *productRepoImlp) GetAllProduct(ctx context.Context) (entities []domain.ProductView, err error) {
	logger.Info(ctx, "GetAllProduct start")
	defer logger.Info(ctx, "GetAllProduct end")

	redisKey := domain.REDIS_KEY_GET_ALL_PRODUCTS
	dur, err := repo.redisCLient.Get(ctx, redisKey, &entities)
	if dur == -2 || err != nil {
		logger.Info(ctx, "redis miss fetching from db")
		sqlQuery := `
        SELECT 
            p."PRODUCT_ID", 
            p."PRODUCT_NAME", 
            p."PRODUCT_DESCRIPTION", 
            p."PRODUCT_IMAGE", 
            p."PRODUCT_PRICE", 
            p."PRODUCT_CATEGORY", 
            p."GENDER", 
            i."INVENTORY_QUANTITY" as "PRODUCT_QUANTITY"
        FROM 
            public."PRODUCT" p
        LEFT JOIN 
            public."INVENTORY" i
        ON 
            p."PRODUCT_ID" = i."PRODUCT_ID";
    `

		rows, err := repo.db.DB.Query(ctx, sqlQuery)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var product domain.ProductView
			err = rows.Scan(
				&product.ProductID,
				&product.ProductName,
				&product.ProductDescription,
				&product.ProductImage,
				&product.ProductPrice,
				&product.ProductCategory,
				&product.Gender,
				&product.ProductQuantity,
			)
			if err != nil {
				return nil, err
			}
			entities = append(entities, product)
		}

		if rows.Err() != nil {
			logger.Errorf(ctx, rows.Err().Error())
			return nil, rows.Err()
		}

		if len(entities) == 0 {
			return nil, fmt.Errorf(errorx.ErrorMessageNoRowAffected)
		}
		go repo.redisCLient.Set(context.Background(), redisKey, entities, 0)
		return entities, nil
	}
	return entities, nil
}

// GetProductByCategory implements domain.IProductRepo.
func (repo *productRepoImlp) GetProductByCategory(ctx context.Context, category string) (entities []domain.ProductView, err error) {
	logger.Info(ctx, "GetProductByGender start")
	defer logger.Info(ctx, "GetProductByGender end")

	redisKey := fmt.Sprintf(domain.REDIS_KEY_CATEGORY+"%s", strings.ToUpper(category))
	dur, err := repo.redisCLient.Get(ctx, redisKey, &entities)
	if dur == -2 || err != nil {
		logger.Info(ctx, "redis miss fetching from db")
		sqlQuery :=
			`
        SELECT 
            p."PRODUCT_ID", 
            p."PRODUCT_NAME", 
            p."PRODUCT_DESCRIPTION", 
            p."PRODUCT_IMAGE", 
            p."PRODUCT_PRICE", 
            p."PRODUCT_CATEGORY", 
            p."GENDER", 
            i."INVENTORY_QUANTITY" as "PRODUCT_QUANTITY"
        FROM 
            public."PRODUCT" p
        LEFT JOIN 
            public."INVENTORY" i
        ON 
            p."PRODUCT_ID" = i."PRODUCT_ID"
        WHERE 
            p."PRODUCT_CATEGORY" = $1;
    `
		rows, err := repo.db.DB.Query(ctx, sqlQuery, category)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var product domain.ProductView
			err = rows.Scan(
				&product.ProductID,
				&product.ProductName,
				&product.ProductDescription,
				&product.ProductImage,
				&product.ProductPrice,
				&product.ProductCategory,
				&product.Gender,
				&product.ProductQuantity,
			)
			if err != nil {
				return nil, err
			}
			entities = append(entities, product)
		}

		if rows.Err() != nil {
			logger.Errorf(ctx, rows.Err().Error())
			return nil, rows.Err()
		}

		if len(entities) == 0 {
			return nil, fmt.Errorf(errorx.ErrorMessageNoRowAffected)
		}
		go repo.redisCLient.Set(context.Background(), redisKey, entities, 0)
		return entities, nil
	}
	return entities, nil
}

// GetProductByGender implements domain.IProductRepo.
func (repo *productRepoImlp) GetProductByGender(ctx context.Context, gender string) (entities []domain.ProductView, err error) {
	logger.Info(ctx, "GetProductByGender start")
	defer logger.Info(ctx, "GetProductByGender end")

	redisKey := fmt.Sprintf(domain.REDIS_KEY_GENDER+"%s", strings.ToUpper(gender))
	dur, err := repo.redisCLient.Get(ctx, redisKey, &entities)
	if dur == -2 || err != nil {
		logger.Info(ctx, "redis miss fetching from db")
		sqlQuery :=
			`
        SELECT 
            p."PRODUCT_ID", 
            p."PRODUCT_NAME", 
            p."PRODUCT_DESCRIPTION", 
            p."PRODUCT_IMAGE", 
            p."PRODUCT_PRICE", 
            p."PRODUCT_CATEGORY", 
            p."GENDER", 
            i."INVENTORY_QUANTITY" as "PRODUCT_QUANTITY"
        FROM 
            public."PRODUCT" p
        LEFT JOIN 
            public."INVENTORY" i
        ON 
            p."PRODUCT_ID" = i."PRODUCT_ID"
        WHERE 
            p."GENDER" = $1;
    `
		rows, err := repo.db.DB.Query(ctx, sqlQuery, gender)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var product domain.ProductView
			err = rows.Scan(
				&product.ProductID,
				&product.ProductName,
				&product.ProductDescription,
				&product.ProductImage,
				&product.ProductPrice,
				&product.ProductCategory,
				&product.Gender,
				&product.ProductQuantity,
			)
			if err != nil {
				return nil, err
			}
			entities = append(entities, product)
		}

		if rows.Err() != nil {
			logger.Errorf(ctx, rows.Err().Error())
			return nil, rows.Err()
		}

		if len(entities) == 0 {
			return nil, fmt.Errorf(errorx.ErrorMessageNoRowAffected)
		}
		go repo.redisCLient.Set(context.Background(), redisKey, entities, 0)
		return entities, nil
	}
	return entities, nil
}

// GetProductByID implements domain.IProductRepo.
func (repo *productRepoImlp) GetProductByID(ctx context.Context, productID string) (entity domain.ProductEntity, err error) {
	logger.Info(ctx, "GetProductByID start")
	defer logger.Info(ctx, "GetProductByID end")

	redisKey := fmt.Sprintf(domain.REDIS_KEY_PRODUCT_ID+"%s", strings.ToUpper(productID))
	dur, err := repo.redisCLient.Get(ctx, redisKey, &entity)
	if dur == -2 || err != nil {
		logger.Info(ctx, "redis miss fetching from db")
		sqlQuery := `
            select * from "PRODUCT" where "PRODUCT_ID"=$1;
        `

		if err = repo.db.DB.Get(ctx, &entity, sqlQuery, productID); err != nil {
			return entity, err
		}
		go repo.redisCLient.Set(context.Background(), redisKey, entity, 0)
		return entity, nil
	}
	return entity, nil
}

// GetProductDetail implements domain.IProductRepo.
func (repo *productRepoImlp) GetProductDetail(ctx context.Context, id string) (entity domain.ProductView, err error) {
	logger.Info(ctx, "GetProductDetail start")
	defer logger.Info(ctx, "GetProductDetail end")

	redisKey := fmt.Sprintf(domain.REDIS_KEY_PRODUCT_DETAIL+"%s", strings.ToUpper(id))
	dur, err := repo.redisCLient.Get(ctx, redisKey, &entity)
	if dur == -2 || err != nil {
		logger.Info(ctx, "redis miss fetching from db")
		sqlQuery := `
        SELECT 
            p."PRODUCT_ID", 
            p."PRODUCT_NAME", 
            p."PRODUCT_DESCRIPTION", 
            p."PRODUCT_IMAGE", 
            p."PRODUCT_PRICE", 
            p."PRODUCT_CATEGORY", 
            p."GENDER", 
            i."INVENTORY_QUANTITY" as "PRODUCT_QUANTITY"
        FROM 
            public."PRODUCT" p
        LEFT JOIN 
            public."INVENTORY" i
        ON 
            p."PRODUCT_ID" = i."PRODUCT_ID"
        WHERE 
            p."PRODUCT_ID" = $1;
    `

		if err = repo.db.DB.Get(ctx, &entity, sqlQuery, id); err != nil {
			return entity, err
		}
		go repo.redisCLient.Set(context.Background(), redisKey, entity, 0)
		return entity, nil
	}
	return entity, nil
}

// Insert implements domain.IProductRepo.
func (repo *productRepoImlp) Insert(ctx context.Context, entity domain.ProductEntity) error {
	panic("unimplemented")
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
