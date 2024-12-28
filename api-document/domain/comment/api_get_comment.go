package domain

import "context"

type GetCommentRequest struct {
}

type GetCommentResponse struct {
}

type GetCommentHandler interface {
	Handle(ctx context.Context, req *GetCommentRequest) (res *GetCommentResponse, err error)
}
