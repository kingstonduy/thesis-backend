package domain

import "context"

// const (
// 	tag = "DeleteCartItem"
// )

type DeleteCartItemRequest struct {
	CartItems []DeleteCartItemDetail `json:"cartItems"`
}

type DeleteCartItemDetail struct {
	UserID string `json:"userID"`
}

type DeleteCartItemResponse struct{}

type DeleteCartItemHandler interface {
	Handle(ctx context.Context, req *DeleteCartItemRequest) (res *DeleteCartItemResponse, err error)
}
