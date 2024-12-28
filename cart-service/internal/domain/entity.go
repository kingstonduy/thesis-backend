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
	AddCartItem(ctx context.Context, params AddCartItemParams) error
	GetCart(ctx context.Context, params GetCartParamsIn) (GetCartParamsOut, error)
	UpdateCartItem(ctx context.Context, params UpdateCartItemParams) error
	DeleteCartItem(ctx context.Context, params DeleteCartItemParams) error
	DeleteUserCart(ctx context.Context, params DeleteUserCartParams) error
}

type DeleteUserCartParams struct {
	UserID string `json:"userId"`
}

type AddCartItemParams struct {
	CartItems []CartItem `json:"cartItems"`
}
type GetCartParamsIn struct {
	UserID string `json:"userId"`
}
type GetCartParamsOut struct {
	CartItems []GetCartItemDetail `json:"cartItems"`
}

type UpdateCartItemParams struct {
	CartItemID       string `json:"cartItemId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type DeleteCartItemParams struct {
	CartItemID string `json:"cartItemId"`
}
