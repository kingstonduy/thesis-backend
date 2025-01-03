package domain

import (
	"context"
)

type RevertTransactionRequest struct {
	UserID  string                           `json:"userID"`
	Details []RevertTransactionRequestDetail `json:"details"`
}

type RevertTransactionRequestDetail struct {
	ProductID        string `json:"productId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type RevertTransactionResponse struct {
}

type IRevertTransactionHandler interface {
	Handle(ctx context.Context, req *RevertTransactionRequest) (*RevertTransactionResponse, error)
}
