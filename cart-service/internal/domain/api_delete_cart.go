package domain

import "context"

// const (
// 	tag = "DeleteCartItem"
// )

type DeleteCartItemsRequest struct {
	Details []DelteCartItemsRequestDetail `json:"details"`
}

type DelteCartItemsRequestDetail struct {
	CartItemID string `json:"cartItemID"`
}
type DeleteCartItemsResponse struct{}

type IDeleteCartItemHandler interface {
	Handle(ctx context.Context, req *DeleteCartItemsRequest) (res *DeleteCartItemsResponse, err error)
}
