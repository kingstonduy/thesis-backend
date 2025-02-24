package domain

import (
	"context"
)

type ReadCartItemEntity struct {
	CartItemID       string `json:"cart_item_id" db:"CART_ITEM_ID"`
	UserID           string `json:"user_id" db:"USER_ID"`
	ProductID        string `json:"product_id" db:"PRODUCT_ID"`
	CartItemQuantity int    `json:"cart_item_quantity" db:"CART_ITEM_QUANTITY"`
}

type IReadCartItemRepo interface {
	Upsert(ctx context.Context, entity ReadCartItemEntity) error
	Delete(ctx context.Context, id string) error
}
