package domain

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type ProductEntity struct {
	ProductID          string          `json:"product_id" db:"PRODUCT_ID"`
	ProductName        string          `json:"product_name" db:"PRODUCT_NAME"`
	ProductDescription string          `json:"product_description" db:"PRODUCT_DESCRIPTION"` // Assuming JSON as a string for ORM
	ProductImage       string          `json:"product_image" db:"PRODUCT_IMAGE"`
	ProductQuantity    int             `json:"product_quantity" db:"PRODUCT_QUANTITY"`
	ProductPrice       decimal.Decimal `json:"product_price" db:"PRODUCT_PRICE"`
	CreatedAt          time.Time       `json:"created_at" db:"CREATED_AT"`
	UpdatedAt          time.Time       `json:"updated_at" db:"UPDATED_AT"`
	ProductCategory    string          `json:"product_category" db:"PRODUCT_CATEGORY"`
	Gender             string          `json:"gender" db:"GENDER"`
}

type IProductRepo interface {
	GetAllProduct(ctx context.Context) ([]ProductEntity, error)
	GetProductByID(ctx context.Context, productID string) (ProductEntity, error)
	UpdateProductByID(ctx context.Context, product ProductEntity) error
}
