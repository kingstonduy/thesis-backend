package domain

import "context"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct{}

type IRLoginHandler interface {
	Handle(ctx context.Context, req *LoginRequest) (res *LoginResponse, err error)
}
