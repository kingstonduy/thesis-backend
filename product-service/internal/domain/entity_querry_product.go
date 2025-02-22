package domain

import (
	"context"
	"time"
)

type ReadProductEntity struct {
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

type IReadProductRepo interface {
	GetProductDetail(ctx context.Context, id string) (entity ReadProductEntity, err error)
	GetAllProduct(ctx context.Context) ([]ReadProductEntity, error)
	GetProductByGender(ctx context.Context, gender string) ([]ReadProductEntity, error)
	GetProductByCategory(ctx context.Context, category string) ([]ReadProductEntity, error)
	GetProductByID(ctx context.Context, productID string) (ReadProductEntity, error)
}
