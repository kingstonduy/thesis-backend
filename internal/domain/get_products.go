package domain

import "context"

type GetProductsRequest struct {
}

type GetProductsResponse struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	ImageURL      string  `json:"imageUrl"`
	Price         float64 `json:"price"`
	AverageRating float64 `json:"averageRating"`
}

type IGetProductsHandler interface {
	Handle(ctx context.Context, req *GetProductsRequest) (res *GetProductsResponse, err error)
}
