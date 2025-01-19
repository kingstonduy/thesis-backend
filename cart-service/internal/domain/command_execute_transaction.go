package domain

import (
	"context"

	"github.com/kingstonduy/go-core/transport"
)

type ExecuteTransactionRequest struct {
	UserID  string                            `json:"userID"`
	Details []ExecuteTransactionRequestDetail `json:"details"`
}

type ExecuteTransactionRequestDetail struct {
	CartItemID       string `json:"cartItemId"`
	ProductID        string `json:"productId"`
	ProductName      string `json:"productName"`
	ProductImage     string `json:"productImage"`
	ProductCatergory string `json:"productCatergory"`
	ProductPrice     string `json:"productPrice"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type ExecuteTransactionResponse struct {
}

type IExecuteTransactionHandler interface {
	Handle(ctx context.Context, req *Command[transport.Request[ExecuteTransactionRequest]]) (*ExecuteTransactionResponse, error)
}
