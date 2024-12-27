package domain

import "context"

type CheckoutRequest struct {
}

type CheckoutResponse struct {
}

type CheckoutHandler interface {
	Handle(ctx context.Context, req *CheckoutRequest) (res *CheckoutResponse, err error)
}
