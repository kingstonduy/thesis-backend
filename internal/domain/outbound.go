package domain

import (
	"context"
	"time"

	"github.com/kingstonduy/go-core/transport"
	"github.com/shopspring/decimal"
)

type GetCommentsByProductIDRequest struct {
	ProductID string `json:"productID"`
}

type GetCommentsByProductIDResponse struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	UserImage string          `json:"userImage"`
	Username  string          `json:"userName"`
	Timestamp time.Time       `json:"timestamp"`
	Content   string          `json:"content"`
	Rating    decimal.Decimal `json:"rating"`
}

type GetAvgerageRatingByProductIDRequest struct {
	ProductID string `json:"productID"`
}
type GetAvgerageRatingByProductIDResponse struct {
	Rating decimal.Decimal `json:"rating"`
}

type ICommentOutbound interface {
	GetAllCommentByProductID(ctx context.Context, req GetCommentsByProductIDRequest, trace transport.Trace) (GetCommentsByProductIDResponse, error)
	GetAvgerageRatingByProductID(ctx context.Context, req GetAvgerageRatingByProductIDRequest, trace transport.Trace) (GetAvgerageRatingByProductIDResponse, error)
}
