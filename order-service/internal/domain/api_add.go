package domain

import "context"

type AddRequest struct {
}

type AddResponse struct {
}

type AddHandler interface {
	Handle(ctx context.Context, req *AddRequest) (res *AddResponse, err error)
}
