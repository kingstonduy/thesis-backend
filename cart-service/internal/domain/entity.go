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
	DeleteCartItemsByID(ctx context.Context, params DeleteCartItemParamsIn) error
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

type DeleteCartItemParamsIn struct {
	Details []DeleteCartItemParamsInDetails
}

type DeleteCartItemParamsInDetails struct {
	CartItemID string `json:"cartItemId"`
}

type ProductCdc struct {
	ProductID          string `json:"PRODUCT_ID" db:"PRODUCT_ID"`
	ProductName        string `json:"PRODUCT_NAME" db:"PRODUCT_NAME"`
	ProductDescription string `json:"PRODUCT_DESCRIPTION" db:"PRODUCT_DESCRIPTION"`
	ProductImage       string `json:"PRODUCT_IMAGE" db:"PRODUCT_IMAGE"`
	ProductQuantity    int    `json:"PRODUCT_QUANTITY" db:"PRODUCT_QUANTITY"`
	ProductPrice       string `json:"PRODUCT_PRICE" db:"PRODUCT_PRICE"`
	CreatedAt          int    `json:"CREATED_AT" db:"CREATED_AT"`
	UpdatedAt          int    `json:"UPDATED_AT" db:"UPDATED_AT"`
	ProductCategory    string `json:"PRODUCT_CATEGORY" db:"PRODUCT_CATEGORY"`
	Gender             string `json:"GENDER" db:"GENDER"`
	AvgRating          string `json:"AVERAGE_RATING" db:"AVERAGE_RATING"`
	TotalRating        int    `json:"TOTAL_RATING" db:"TOTAL_RATING"`
}

type IProductRepo interface {
	Insert(ctx context.Context, entity ProductCdc) error
	Update(ctx context.Context, cols map[string]interface{}, condition map[string]interface{}) error
}
