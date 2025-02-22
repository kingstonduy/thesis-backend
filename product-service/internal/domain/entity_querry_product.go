package domain

import (
	"context"
	"time"
)

type ReadProductEntity struct {
	ProductID          string    `bson:"PRODUCT_ID" json:"PRODUCT_ID"`
	ProductName        string    `bson:"PRODUCT_NAME" json:"PRODUCT_NAME"`
	ProductDescription string    `bson:"PRODUCT_DESCRIPTION" json:"PRODUCT_DESCRIPTION"`
	ProductImage       string    `bson:"PRODUCT_IMAGE" json:"PRODUCT_IMAGE"`
	ProductPrice       string    `bson:"PRODUCT_PRICE" json:"PRODUCT_PRICE"`
	ProductQuantity    int       `bson:"INVENTORY_QUANTITY" json:"PRODUCT_QUANTITY"`
	ProductCategory    string    `bson:"PRODUCT_CATEGORY" json:"PRODUCT_CATEGORY"`
	Gender             string    `bson:"GENDER" json:"GENDER"`
	CreatedAt          time.Time `bson:"_insertedTS" json:"CREATED_AT"`
	UpdatedAt          time.Time `bson:"_modifiedTS" json:"UPDATED_AT"`
}

type IReadProductRepo interface {
	GetProductDetail(ctx context.Context, id string) (entity ReadProductEntity, err error)
	GetAllProduct(ctx context.Context) ([]ReadProductEntity, error)
	GetProductByGender(ctx context.Context, gender string) ([]ReadProductEntity, error)
	GetProductByCategory(ctx context.Context, category string) ([]ReadProductEntity, error)
	GetProductByID(ctx context.Context, productID string) (ReadProductEntity, error)
}
