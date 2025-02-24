package domain

import (
	"context"
)

type ReadCartEntity struct {
	UserID string           `bson:"USER_ID" json:"user_id"`
	Detail []ReadCartDetail `bson:"CART_ITEMS" json:"cart_items"`
}

type ReadCartDetail struct {
	CartItemID       string `json:"cart_item_id" bson:"CART_ITEM_ID"`
	ProductID        string `json:"product_id" bson:"PRODUCT_ID"`
	ProductName      string `json:"product_name" bson:"PRODUCT_NAME"`
	ProductImage     string `bson:"PRODUCT_IMAGE" json:"PRODUCT_IMAGE"`
	ProductCategory  string `bson:"PRODUCT_CATEGORY" json:"PRODUCT_CATEGORY"`
	ProductPrice     string `bson:"PRODUCT_PRICE" json:"PRODUCT_PRICE"`
	Gender           string `bson:"GENDER" json:"GENDER"`
	CartItemQuantity int    `json:"cart_item_quantity" bson:"CART_ITEM_QUANTITY"`
}

type IReadCartItemRepo interface {
	Upsert(ctx context.Context, entity ReadCartEntity) error
	Delete(ctx context.Context, id string) error
	GetCartItemByUserID(ctx context.Context, userID string) (ReadCartEntity, error)
}
