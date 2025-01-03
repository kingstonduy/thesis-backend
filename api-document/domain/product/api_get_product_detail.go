package domain

import (
	"context"
)

type GetProductDetailRequest struct {
	ID string `json:"productId"`
}

type GetProductDetailResponse struct {
	ID              string `json:"productId"`
	Name            string `json:"productName"`
	Catergory       string `json:"productCatergory"`
	Price           string `json:"productPrice"`
	Description     string `json:"productDescription"`
	Image           string `json:"productImage"`
	ProductQuantity int    `json:"productQuantity"`
	AverageRating   string `json:"averageRating"`
	TotalRating     int    `json:"totalRating"`
}

type IGetProductDetailHandler interface {
	Handle(ctx context.Context, req *GetProductDetailRequest) (res *GetProductDetailResponse, err error)
}
