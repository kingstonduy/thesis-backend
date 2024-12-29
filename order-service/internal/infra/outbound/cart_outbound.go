package outbound

import (
	"context"

	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/order-service/internal/domain"
)

type cartOutbound struct{}

func NewCartOutbound() domain.ICartOutbound {
	return &cartOutbound{}
}

// ExecuteTransaction implements domain.ICartOutbound.
func (c *cartOutbound) ExecuteTransaction(ctx context.Context, reqType transport.Request[domain.CartExecuteTransactionRequest]) (resType transport.Response[domain.CartExecuteTransactionResponse], err error) {
	panic("unimplemented")
}

// RevertTransaction implements domain.ICartOutbound.
func (c *cartOutbound) RevertTransaction(ctx context.Context, reqType transport.Request[domain.CartExecuteTransactionRequest]) (resType transport.Response[domain.CartExecuteTransactionResponse], err error) {
	panic("unimplemented")
}
