package domain

import "context"

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
	CartItemID       string `json:"cartItemId"`
	ProductID        string `json:"productId"`
	ProductName      string `json:"productName"`
	ProductImage     string `json:"productImage"`
	ProductCatergory string `json:"productCatergory"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type GetCartHandler interface {
	Handle(ctx context.Context, req *GetCartRequest) (res *GetCartResponse, err error)
}
