package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

type GetProductsRequest struct {
}

type GetProductsResponse struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	ImageURL      string          `json:"imageUrl"`
	Price         decimal.Decimal `json:"price"`
	AverageRating decimal.Decimal `json:"averageRating"`
}

type IGetProductsHandler interface {
	Handle(ctx context.Context, req *GetProductsRequest) (res *GetProductsResponse, err error)
}
