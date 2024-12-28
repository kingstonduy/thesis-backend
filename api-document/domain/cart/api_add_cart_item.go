package domain

import "context"

// const (
//     tag = "AddCartItem"
// )

type AddCartItemRequest struct {
	UserID    string              `json:"userId"`
	CartItems []AddCartItemDetail `json:"cartItems"`
}

type AddCartItemResponse struct {
}

type AddCartItemDetail struct {
	ProductID        string `json:"productId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type AddCartItemHandler interface {
	Handle(ctx context.Context, req *AddCartItemRequest) (res *AddCartItemResponse, err error)
}
