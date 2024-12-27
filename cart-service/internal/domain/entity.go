package domain

import (
	"context"
	"database/sql"
)

type CartItem struct {
	CartItemID       string       `json:"cart_item_id" db:"CART_ITEM_ID"`
	UserID           string       `json:"user_id" db:"USER_ID"`
	ProductID        string       `json:"product_id" db:"PRODUCT_ID"`
	CartItemQuantity float64      `json:"cart_item_quantity" db:"CART_ITEM_QUANTITY"`
	ProductName      string       `json:"product_name" db:"PRODUCT_NAME"`
	ProductImage     string       `json:"product_image" db:"PRODUCT_IMAGE"`
	ProductPrice     float64      `json:"product_price" db:"PRODUCT_PRICE"`
	ProductCategory  string       `json:"product_category" db:"PRODUCT_CATEGORY"`
	CreatedAt        sql.NullTime `json:"created_at" db:"CREATED_AT"`
	UpdatedAt        sql.NullTime `json:"updated_at" db:"UPDATED_AT"`
}

type IProductRepo interface {
	DeleteByUserID(ctx context.Context, userID string) error
	GetCartItemsByUserID(ctx context.Context, userID string) ([]CartItem, error)
	UpdateCartItem(ctx context.Context, cartItemID string) error
}
