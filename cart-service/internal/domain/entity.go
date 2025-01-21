package domain

import (
	"context"
	"time"
)

type CartItem struct {
	CartItemID       string    `json:"cart_item_id" db:"CART_ITEM_ID"`
	UserID           string    `json:"user_id" db:"USER_ID"`
	ProductID        string    `json:"product_id" db:"PRODUCT_ID"`
	CartItemQuantity int       `json:"cart_item_quantity" db:"CART_ITEM_QUANTITY"`
	CreatedAt        time.Time `json:"created_at" db:"CREATED_AT"`
	UpdatedAt        time.Time `json:"updated_at" db:"UPDATED_AT"`
}

type ICartRepo interface {
	Insert(ctx context.Context, entity CartItem) error
	Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error
	DeleteById(ctx context.Context, id string) error
	// VIEW
	GetCart(ctx context.Context, params GetCartParamsIn) (GetCartParamsOut, error)
}

type GetCartParamsIn struct {
	UserID string `json:"userId"`
}
type GetCartParamsOut struct {
	CartItems []GetCartItemDetail `json:"cartItems"`
}
