package domain

import (
	"context"

	cmd_pipeline "github.com/kingstonduy/go-core/comman-pipeline"
)

type RevertTransactionRequest struct {
	UserID  string                           `json:"userID"`
	Details []RevertTransactionRequestDetail `json:"details"`
}

type RevertTransactionRequestDetail struct {
	CartItemID       string `json:"cartItemId"`
	ProductID        string `json:"productId"`
	ProductName      string `json:"productName"`
	ProductImage     string `json:"productImage"`
	ProductCatergory string `json:"productCatergory"`
	ProductPrice     string `json:"productPrice"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type RevertTransactionResponse struct {
}

type IRevertTransactionHandler interface {
	Handle1(ctx context.Context, outbox cmd_pipeline.OutboxWithTrace) error
}
