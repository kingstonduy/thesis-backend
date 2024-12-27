package domain

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type CartItem struct {
	CartItemID       string          `json:"cart_item_id" db:"CART_ITEM_ID"`
	UserID           string          `json:"user_id" db:"USER_ID"`
	ProductID        string          `json:"product_id" db:"PRODUCT_ID"`
	CartItemQuantity int             `json:"cart_item_quantity" db:"CART_ITEM_QUANTITY"`
	ProductName      string          `json:"product_name" db:"PRODUCT_NAME"`
	ProductImage     string          `json:"product_image" db:"PRODUCT_IMAGE"`
	ProductPrice     decimal.Decimal `json:"product_price" db:"PRODUCT_PRICE"`
	ProductCategory  string          `json:"product_category" db:"PRODUCT_CATEGORY"`
	CreatedAt        time.Time       `json:"created_at" db:"CREATED_AT"`
	UpdatedAt        time.Time       `json:"updated_at" db:"UPDATED_AT"`
}

type ICartRepo interface {
	AddCartItem(ctx context.Context, params AddCartItemParams) error
	GetCart(ctx context.Context, params GetCartParams) ([]CartItem, error)
	UpdateCartItem(ctx context.Context, params UpdateCartItemParams) error
	DeleteCartItem(ctx context.Context, params DeleteCartItemParams) error
}

type AddCartItemParams struct {
	CartItems []CartItem `json:"cartItems"`
}
type GetCartParams struct {
	UserID string `json:"userId"`
}

type UpdateCartItemParams struct {
	CartItemID       string `json:"cartItemId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type DeleteCartItemParams struct {
	UserID string `json:"userID"`
}
