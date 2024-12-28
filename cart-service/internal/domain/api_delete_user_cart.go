package domain

import "context"

type DeleteUserCartRequest struct {
	UserID string `json:"userId"`
}
type DeleteUserCartResponse struct{}

type DeleteUserCartHandler interface {
	Handle(ctx context.Context, req *DeleteUserCartRequest) (res *DeleteUserCartResponse, err error)
}
