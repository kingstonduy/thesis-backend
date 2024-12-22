package configuration

import (
	"github.com/go-resty/resty/v2"
	hresty "github.com/kingstonduy/go-core/client/http/resty"
)

func NewRestyClient() *resty.Client {
	client := hresty.NewRestyClient()

	return client
}
