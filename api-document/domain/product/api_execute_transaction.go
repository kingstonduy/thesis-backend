package domain

import (
	"context"
)

type ExecuteTransactionRequest struct {
	UserID  string                            `json:"userID"`
	Details []ExecuteTransactionRequestDetail `json:"details"`
}

type ExecuteTransactionRequestDetail struct {
	ProductID        string `json:"productId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type ExecuteTransactionResponse struct {
}

type IExecuteTransactionHandler interface {
	Handle(ctx context.Context, req *ExecuteTransactionRequest) (*ExecuteTransactionResponse, error)
}
