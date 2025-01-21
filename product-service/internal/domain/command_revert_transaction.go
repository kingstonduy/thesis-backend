package domain

import (
	"context"

	"github.com/kingstonduy/go-core/transport"
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
	Handle(ctx context.Context, cmd *Command[transport.Request[RevertTransactionRequest]]) (*RevertTransactionResponse, error)
}
