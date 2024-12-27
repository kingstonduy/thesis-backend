package domain

import "context"

// const (
//     tag = "AddCartItem"
// )

type AddCartItemRequest struct {
	CartItems []AddCartItemDetail `json:"cartItems"`
}

type AddCartItemDetail struct {
	CartItemID       string `json:"cartItemId"`
	ProductID        string `json:"productId"`
	ProductName      string `json:"productName"`
	ProductImage     string `json:"productImage"`
	ProductCatergory string `json:"productCatergory"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}
type AddCartItemResponse struct {
}

type AddCartItemHandler interface {
	Handle(ctx context.Context, req *AddCartItemRequest) (res *AddCartItemResponse, err error)
}
