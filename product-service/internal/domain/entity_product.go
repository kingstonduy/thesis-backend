package domain

import (
	"context"
	"time"
)

const (
	REDIS_KEY_GET_ALL_PRODUCTS = "PRODUCTS_ALL"
	REDIS_KEY_PRODUCT_ID       = "PRODUCT_ID-"
	REDIS_KEY_GENDER           = "PRODUCT_GENDER_"
	REDIS_KEY_CATEGORY         = "PRODUCT_CATEGORY_"
	REDIS_KEY_PRODUCT_DETAIL   = "PRODUCT_DETAIL-"
)

type ProductEntity struct {
	ProductID          string `json:"PRODUCT_ID" db:"PRODUCT_ID"`
	ProductName        string `json:"PRODUCT_NAME" db:"PRODUCT_NAME"`
	ProductDescription string `json:"PRODUCT_DESCRIPTION" db:"PRODUCT_DESCRIPTION"`
	ProductImage       string `json:"PRODUCT_IMAGE" db:"PRODUCT_IMAGE"`
	ProductPrice       string `json:"PRODUCT_PRICE" db:"PRODUCT_PRICE"`
	ProductCategory    string `json:"PRODUCT_CATEGORY" db:"PRODUCT_CATEGORY"`
	Gender             string `json:"GENDER" db:"GENDER"`
	// CreatedAt          time.Time `json:"CREATED_AT" db:"CREATED_AT"`
	// UpdatedAt          time.Time `json:"UPDATED_AT" db:"UPDATED_AT"`
}

type ProductView struct {
	ProductID          string    `json:"PRODUCT_ID" db:"PRODUCT_ID"`
	ProductName        string    `json:"PRODUCT_NAME" db:"PRODUCT_NAME"`
	ProductDescription string    `json:"PRODUCT_DESCRIPTION" db:"PRODUCT_DESCRIPTION"`
	ProductImage       string    `json:"PRODUCT_IMAGE" db:"PRODUCT_IMAGE"`
	ProductPrice       string    `json:"PRODUCT_PRICE" db:"PRODUCT_PRICE"`
	ProductQuantity    int       `json:"PRODUCT_QUANTITY" db:"PRODUCT_QUANTITY"`
	ProductCategory    string    `json:"PRODUCT_CATEGORY" db:"PRODUCT_CATEGORY"`
	Gender             string    `json:"GENDER" db:"GENDER"`
	CreatedAt          time.Time `json:"CREATED_AT" db:"CREATED_AT"`
	UpdatedAt          time.Time `json:"UPDATED_AT" db:"UPDATED_AT"`
}

type IProductRepo interface {
	GetProductDetail(ctx context.Context, id string) (entity ProductView, err error)
	GetAllProduct(ctx context.Context) ([]ProductView, error)
	GetProductByGender(ctx context.Context, gender string) ([]ProductView, error)
	GetProductByCategory(ctx context.Context, category string) ([]ProductView, error)

	GetProductByID(ctx context.Context, productID string) (ProductEntity, error)
	Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error
	Insert(ctx context.Context, entity ProductEntity) error
}
