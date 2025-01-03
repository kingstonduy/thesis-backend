package domain

import (
	"context"
	"time"
)

type ProductEntity struct {
	ProductID          string    `json:"PRODUCT_ID" db:"PRODUCT_ID"`
	ProductName        string    `json:"PRODUCT_NAME" db:"PRODUCT_NAME"`
	ProductDescription string    `json:"PRODUCT_DESCRIPTION" db:"PRODUCT_DESCRIPTION"`
	ProductImage       string    `json:"PRODUCT_IMAGE" db:"PRODUCT_IMAGE"`
	ProductQuantity    int       `json:"PRODUCT_QUANTITY" db:"PRODUCT_QUANTITY"`
	ProductPrice       string    `json:"PRODUCT_PRICE" db:"PRODUCT_PRICE"`
	CreatedAt          time.Time `json:"CREATED_AT" db:"CREATED_AT"`
	UpdatedAt          time.Time `json:"UPDATED_AT" db:"UPDATED_AT"`
	ProductCategory    string    `json:"PRODUCT_CATEGORY" db:"PRODUCT_CATEGORY"`
	Gender             string    `json:"GENDER" db:"GENDER"`
	AvgRating          string    `json:"AVERAGE_RATING" db:"AVERAGE_RATING"`
	TotalRating        int       `json:"TOTAL_RATING" db:"TOTAL_RATING"`
}

type IProductRepo interface {
	GetAllProduct(ctx context.Context) ([]ProductEntity, error)
	GetProductByID(ctx context.Context, productID string) (ProductEntity, error)
	UpdateProductByID(ctx context.Context, product ProductEntity) error
	GetProductByGender(ctx context.Context, gender string) ([]ProductEntity, error)
	GetProductByCategory(ctx context.Context, category string) ([]ProductEntity, error)
}

type ProductCdc struct {
	ProductID          string `json:"PRODUCT_ID" db:"PRODUCT_ID"`
	ProductName        string `json:"PRODUCT_NAME" db:"PRODUCT_NAME"`
	ProductDescription string `json:"PRODUCT_DESCRIPTION" db:"PRODUCT_DESCRIPTION"`
	ProductImage       string `json:"PRODUCT_IMAGE" db:"PRODUCT_IMAGE"`
	ProductQuantity    int    `json:"PRODUCT_QUANTITY" db:"PRODUCT_QUANTITY"`
	ProductPrice       string `json:"PRODUCT_PRICE" db:"PRODUCT_PRICE"`
	CreatedAt          int    `json:"CREATED_AT" db:"CREATED_AT"`
	UpdatedAt          int    `json:"UPDATED_AT" db:"UPDATED_AT"`
	ProductCategory    string `json:"PRODUCT_CATEGORY" db:"PRODUCT_CATEGORY"`
	Gender             string `json:"GENDER" db:"GENDER"`
	AvgRating          string `json:"AVERAGE_RATING" db:"AVERAGE_RATING"`
	TotalRating        int    `json:"TOTAL_RATING" db:"TOTAL_RATING"`
}
