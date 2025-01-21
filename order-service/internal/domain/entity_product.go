package domain

import "context"

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

type IProductRepo interface {
	Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error
	Insert(ctx context.Context, entity ProductEntity) error
}
