package outbound

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/kingstonduy/go-core/transport"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
)

type cartOutbound struct {
	httpConfig  *configuration.HttpConfig
	restyClient *resty.Client
}

func NewCartOutbound(
	restyClient *resty.Client,
	httpConfig *configuration.HttpConfig,
) domain.ICartOutbound {
	return &cartOutbound{
		httpConfig:  httpConfig,
		restyClient: restyClient,
	}
}

// ExecuteTransaction implements domain.ICartOutbound.
func (o *cartOutbound) ExecuteTransaction(ctx context.Context, reqType transport.Request[domain.CartExecuteTransactionRequest]) (resType transport.Response[domain.CartExecuteTransactionResponse], err error) {
	// TODO remove hard code
	baseURL := o.httpConfig.BaseUrl
	path := o.httpConfig.ExecuteTransactionCartUrl
	fullURL := baseURL + path

	headers := transport.GetHttpRequestHeaderByCtx(ctx)
	resp, err := o.restyClient.R().SetContext(ctx).SetHeaders(headers).SetBody(reqType).SetResult(&resType).Post(fullURL)
	if err != nil {
		return resType, err
	}

	if resp.StatusCode() != 200 {
		return resType, fmt.Errorf(resp.String())
	}

	return resType, nil
}

// RevertTransaction implements domain.ICartOutbound.
func (o *cartOutbound) RevertTransaction(ctx context.Context, reqType transport.Request[domain.CartExecuteTransactionRequest]) (resType transport.Response[domain.CartExecuteTransactionResponse], err error) {
	panic("unimplemented")
}
