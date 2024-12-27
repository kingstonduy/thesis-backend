package domain

import "context"

type GetCheckoutItemRequest struct {
}

type GetCheckoutItemResponse struct {
}

type GetCheckoutItemHandler interface {
	Handle(ctx context.Context, req *GetCheckoutItemRequest) (res *GetCheckoutItemResponse, err error)
}
