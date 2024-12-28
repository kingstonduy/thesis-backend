package domain

import "context"

// const (
// 	tag = "DeleteCartItem"
// )

type DeleteCartItemRequest struct {
	CartItemID string `json:"cartItemID"`
}

type DeleteCartItemResponse struct{}

type DeleteCartItemHandler interface {
	Handle(ctx context.Context, req *DeleteCartItemRequest) (res *DeleteCartItemResponse, err error)
}
