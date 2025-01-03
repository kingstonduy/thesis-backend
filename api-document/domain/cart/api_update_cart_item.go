package domain

import "context"

// const (
//     tag = "UpdateCartItem"
// )

type UpdateCartItemRequest struct {
	CartItemID       string `json:"cartItemId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type UpdateCartItemResponse struct {
}

type IUpdateCartItemHandler interface {
	Handle(ctx context.Context, req *UpdateCartItemRequest) (res *UpdateCartItemResponse, err error)
}
