package outbound

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/kingstonduy/go-core/transport"
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
)

type productOutbound struct {
	httpConfig  *configuration.HttpConfig
	restyClient *resty.Client
}

func NewProductOutbound(
	httpConfig *configuration.HttpConfig,
	restyClient *resty.Client,
) domain.IProductOutbound {
	return &productOutbound{
		httpConfig:  httpConfig,
		restyClient: restyClient,
	}
}

// ExecuteTransaction implements domain.IProductOutbound.
func (o *productOutbound) ExecuteTransaction(ctx context.Context, reqType transport.Request[domain.ProductExecuteTransactionRequest]) (resType transport.Response[domain.ProductExecuteTransactionResponse], err error) {
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

// RevertTransaction implements domain.IProductOutbound.
func (o *productOutbound) RevertTransaction(ctx context.Context, reqType transport.Request[domain.ProductExecuteTransactionRequest]) (resType transport.Response[domain.ProductExecuteTransactionResponse], err error) {
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
