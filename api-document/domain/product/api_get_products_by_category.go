package domain

import "context"

type GetProductsByCategoryRequest struct {
	Category string `json:"Category"`
}

type GetProductsByCategoryResponse struct {
	Details []GetProductsByCategoryResponseDetail `json:"details"`
}

type GetProductsByCategoryResponseDetail struct {
	ID              string `json:"productId"`
	Name            string `json:"productName"`
	ImageURL        string `json:"productImage"`
	Price           string `json:"productPrice"`
	AverageRating   string `json:"averageRating"`
	ProductQuantity int    `json:"productQuantity"`
	TotalRating     int    `json:"totalRating"`
}

type IGetProductsByCategoryHandler interface {
	Handle(ctx context.Context, req *GetProductsByCategoryRequest) (res *GetProductsByCategoryResponse, err error)
}
