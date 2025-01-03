package domain

import (
	"context"
)

type GetAllProductRequest struct {
}

type GetAllProductResponse struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID              string `json:"productId"`
	Name            string `json:"productName"`
	ImageURL        string `json:"productImage"`
	Price           string `json:"productPrice"`
	AverageRating   string `json:"averageRating"`
	ProductQuantity int    `json:"productQuantity"`
	TotalRating     int    `json:"totalRating"`
}

type IGetProductsHandler interface {
	Handle(ctx context.Context, req *GetAllProductRequest) (res *GetAllProductResponse, err error)
}
