package domain

import "context"

type ReadProductEntity struct {
	ProductID          string `bson:"PRODUCT_ID" json:"PRODUCT_ID"`
	ProductName        string `bson:"PRODUCT_NAME" json:"PRODUCT_NAME"`
	ProductDescription string `bson:"PRODUCT_DESCRIPTION" json:"PRODUCT_DESCRIPTION"`
	ProductImage       string `bson:"PRODUCT_IMAGE" json:"PRODUCT_IMAGE"`
	ProductPrice       string `bson:"PRODUCT_PRICE" json:"PRODUCT_PRICE"`
	ProductQuantity    int    `bson:"INVENTORY_QUANTITY" json:"PRODUCT_QUANTITY"`
	ProductCategory    string `bson:"PRODUCT_CATEGORY" json:"PRODUCT_CATEGORY"`
	Gender             string `bson:"GENDER" json:"GENDER"`
}

type IReadProductRepo interface {
	GetProductByID(ctx context.Context, productID string) (ReadProductEntity, error)
}
