package outbound

import (
	"context"

	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/order-service/internal/domain"
)

type productOutbound struct{}

func NewProductOutbound() domain.IProductOutbound {
	return &productOutbound{}
}

// ExecuteTransaction implements domain.IProductOutbound.
func (p *productOutbound) ExecuteTransaction(ctx context.Context, reqType transport.Request[domain.ProductExecuteTransactionRequest]) (resType transport.Response[domain.ProductExecuteTransactionResponse], err error) {
	panic("unimplemented")
}

// RevertTransaction implements domain.IProductOutbound.
func (p *productOutbound) RevertTransaction(ctx context.Context, reqType transport.Request[domain.ProductExecuteTransactionRequest]) (resType transport.Response[domain.ProductExecuteTransactionResponse], err error) {
	panic("unimplemented")
}
