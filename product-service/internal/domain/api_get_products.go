package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

type GetAllProductRequest struct {
}

type GetAllProductResponse struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID              string          `json:"productId"`
	Name            string          `json:"productName"`
	ImageURL        string          `json:"productImage"`
	Price           decimal.Decimal `json:"productPrice"`
	AverageRating   decimal.Decimal `json:"averageRating"`
	ProductQuantity int             `json:"productQuantity"`
}

type IGetProductsHandler interface {
	Handle(ctx context.Context, req *GetAllProductRequest) (res *GetAllProductResponse, err error)
}
