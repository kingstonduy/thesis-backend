package domain

import (
	"context"
	"time"
)

type CommentEntity struct {
	CommentID        string    `json:"comment_id" db:"COMMENT_ID"`
	ProductID        string    `json:"product_id" db:"PRODUCT_ID"`
	UserID           string    `json:"user_id" db:"USER_ID"`
	ReviewEvaluation int       `json:"review_evaluation" db:"REVIEW_EVALUATION"`
	ReviewDetail     string    `json:"review_detail" db:"REVIEW_DETAIL"`
	CreatedAt        time.Time `json:"created_at" db:"CREATED_AT"`
	UpdatedAt        time.Time `json:"updated_at" db:"UPDATED_AT"`
}

type ICommentRepo interface {
	GetCommennt(ctx context.Context, params GetCommentParamsIn) (GetCommentParamsOut, error)
	AddComment(ctx context.Context, params AddCommentParams) error
}

type GetCommentParamsIn struct {
	ProductID string `json:"product_id" db:"PRODUCT_ID"`
	Filter    int    `json:"filter" db:"FILTER"`
	Sort      string `json:"sort" db:"SORT"`
}
type GetCommentParamsOut struct {
	Comments []*GetCommentParamsDetail
}

type GetCommentParamsDetail struct {
	CommentID        string    `json:"comment_id" db:"COMMENT_ID"`
	ReviewEvaluation int       `json:"review_evaluation" db:"REVIEW_EVALUATION"`
	ReviewDetail     string    `json:"review_detail" db:"REVIEW_DETAIL"`
	CreatedAt        time.Time `json:"created_at" db:"CREATED_AT"`
	UserID           string    `json:"user_id" db:"USER_ID"`
	UserName         string    `json:"user_name" db:"USER_NAME"`
	UserImage        string    `json:"user_image" db:"USER_IMAGE"`
}

type AddCommentParams struct {
	ProductID        string `json:"product_id" db:"PRODUCT_ID"`
	UserID           string `json:"user_id" db:"USER_ID"`
	ReviewEvaluation int    `json:"review_evaluation" db:"REVIEW_EVALUATION"`
	ReviewDetail     string `json:"review_detail" db:"REVIEW_DETAIL"`
}
