package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

// const (
// 	tag = "GetCart"
// )

type GetCartRequest struct {
	UserID string `json:"userId"`
}
type GetCartResponse struct {
	CartItems []GetCartItemDetail `json:"cartItems"`
}
type GetCartItemDetail struct {
	CartItemID       string          `json:"cartItemId"`
	ProductID        string          `json:"productId"`
	ProductName      string          `json:"productName"`
	ProductImage     string          `json:"productImage"`
	ProductCatergory string          `json:"productCatergory"`
	ProductPrice     decimal.Decimal `json:"productPrice"`
	CartItemQuantity int             `json:"cartItemQuantity"`
}

type IGetCartHandler interface {
	Handle(ctx context.Context, req *GetCartRequest) (res *GetCartResponse, err error)
}
