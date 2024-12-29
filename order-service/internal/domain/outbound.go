package domain

import (
	"context"

	"github.com/kingstonduy/go-core/transport"
)

type CartExecuteTransactionRequest struct {
	Details []CartExecuteTransactionRequestDetails `json:"details,omitempty"`
}
type CartExecuteTransactionResponse struct{}
type CartExecuteTransactionRequestDetails struct {
	CartItemID string `json:"cartItemID,omitempty"`
}

type CartRevertTransactionRequest struct {
	UserID  string                                `json:"userId"`
	Details []CartRevertTransactionRequestDetails `json:"details"`
}
type CartRevertTransactionResponse struct{}
type CartRevertTransactionRequestDetails struct {
	ProductID        string `json:"productId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type ICartOutbound interface {
	ExecuteTransaction(ctx context.Context, reqType transport.Request[CartExecuteTransactionRequest]) (resType transport.Response[CartExecuteTransactionResponse], err error)
	// goi api add cart item
	RevertTransaction(ctx context.Context, reqType transport.Request[CartExecuteTransactionRequest]) (resType transport.Response[CartExecuteTransactionResponse], err error)
}

type ProductExecuteTransactionRequest struct {
	Details []ProductExecuteTransactionRequestDetails `json:"details,omitempty"`
}
type ProductExecuteTransactionResponse struct{}
type ProductExecuteTransactionRequestDetails struct {
	ProductID       string `json:"productId"`
	ProductQuantity int    `json:"productQuantity"`
}

type ProductRevertTransactionRequest struct {
	Details []ProductRevertTransactionRequestDetails `json:"details,omitempty"`
}
type ProductRevertTransactionResponse struct{}
type ProductRevertTransactionRequestDetails struct {
	ProductID       string `json:"productId"`
	ProductQuantity int    `json:"productQuantity"`
}

type IProductOutbound interface {
	ExecuteTransaction(ctx context.Context, reqType transport.Request[ProductExecuteTransactionRequest]) (resType transport.Response[ProductExecuteTransactionResponse], err error)
	RevertTransaction(ctx context.Context, reqType transport.Request[ProductExecuteTransactionRequest]) (resType transport.Response[ProductExecuteTransactionResponse], err error)
}
