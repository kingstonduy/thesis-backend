package domain

import "context"

// const (
//     tag = "AddCartItem"
// )

type AddCartItemRequest struct {
	UserID           string `json:"userId"`
	ProductID        string `json:"productId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type AddCartItemResponse struct {
}

type IAddCartItemHandler interface {
	Handle(ctx context.Context, req *AddCartItemRequest) (res *AddCartItemResponse, err error)
}
