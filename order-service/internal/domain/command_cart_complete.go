package domain

import (
	"context"

	cmd_pipeline "github.com/kingstonduy/go-core/comman-pipeline"
)

type CartCompletedRequest struct {
	UserID  string                       `json:"userID"`
	Details []CartCompletedRequestDetail `json:"details"`
}

type CartCompletedRequestDetail struct {
	CartItemID       string `json:"cartItemId"`
	ProductID        string `json:"productId"`
	ProductName      string `json:"productName"`
	ProductImage     string `json:"productImage"`
	ProductCatergory string `json:"productCatergory"`
	ProductPrice     string `json:"productPrice"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type CartCompletedResponse struct{}

type ICartCompletedHandler interface {
	Handle1(ctx context.Context, outbox cmd_pipeline.OutboxWithTrace) error
}
