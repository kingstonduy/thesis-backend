package outbound

import (
	"context"
	"math/rand"
	"time"

	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/thesis-backend/internal/domain"
	"github.com/shopspring/decimal"
)

type CommentOutbound struct{}

func NewCommentOutbound() domain.ICommentOutbound {
	return &CommentOutbound{}
}

// GetAllCommentByProductID implements domain.ICommentOutbound.
func (c *CommentOutbound) GetAllCommentByProductID(ctx context.Context, req domain.GetCommentsByProductIDRequest, trace transport.Trace) (domain.GetCommentsByProductIDResponse, error) {
	// todo replace later
	return domain.GetCommentsByProductIDResponse{
		Comments: []domain.Comment{
			{
				UserImage: "https://logolook.net/wp-content/uploads/2022/12/GitHub-Logo.png",
				Username:  "kingstonduy",
				Timestamp: time.Now(),
				Content:   "hello it is a test comment",
				Rating:    decimal.NewFromInt(4),
			},
			{
				UserImage: "https://logolook.net/wp-content/uploads/2022/12/GitHub-Logo.png",
				Username:  "kingstonduy1",
				Timestamp: time.Now(),
				Content:   "hello it is a test comment",
				Rating:    decimal.NewFromInt(4),
			},
			{
				UserImage: "https://logolook.net/wp-content/uploads/2022/12/GitHub-Logo.png",
				Username:  "kingstonduy2",
				Timestamp: time.Now(),
				Content:   "hello it is a test comment",
				Rating:    decimal.NewFromInt(4),
			},
		},
	}, nil
}

// GetAvgerageRatingByProductID implements domain.ICommentOutbound.
func (c *CommentOutbound) GetAvgerageRatingByProductID(ctx context.Context, req domain.GetAvgerageRatingByProductIDRequest, trace transport.Trace) (res domain.GetAvgerageRatingByProductIDResponse, err error) {
	// todo replace later

	// generate a random float
	avgRating := rand.Float64() * 5.0

	return domain.GetAvgerageRatingByProductIDResponse{
		Rating: decimal.NewFromFloat(avgRating),
	}, nil
}
