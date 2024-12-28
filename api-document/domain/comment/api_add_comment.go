package domain

import "context"

type AddCommentRequest struct {
	UserID    string `json:"userId,omitempty" validate:"required"`
	ProductID string `json:"productId,omitempty" validate:"required"`
	Content   string `json:"content,omitempty" validate:"required"`
	Rating    int    `json:"rating,omitempty" validate:"required,oneof=1 2 3 4 5"`
}

type AddCommentResponse struct {
}

type IAddCommentHandler interface {
	Handle(ctx context.Context, req *AddCommentRequest) (res *AddCommentResponse, err error)
}
