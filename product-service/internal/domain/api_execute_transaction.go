package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

type ExecuteTransactionRequest struct {
	UserID  string                            `json:"userID"`
	Details []ExecuteTransactionRequestDetail `json:"details"`
}

type ExecuteTransactionRequestDetail struct {
	CartItemID       string          `json:"cartItemId"`
	ProductID        string          `json:"productId"`
	ProductName      string          `json:"productName"`
	ProductImage     string          `json:"productImage"`
	ProductCatergory string          `json:"productCatergory"`
	ProductPrice     decimal.Decimal `json:"productPrice"`
	CartItemQuantity int             `json:"cartItemQuantity"`
}

type ExecuteTransactionResponse struct {
}

type IExecuteTransactionHandler interface {
	Handle(ctx context.Context, req *ExecuteTransactionRequest) (*ExecuteTransactionResponse, error)
}
