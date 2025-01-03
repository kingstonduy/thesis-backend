package domain

import "context"

type GetProductsByGenderRequest struct {
	Gender string `json:"gender"`
}

type GetProductsByGenderResponse struct {
	Details []GetProductsByGenderResponseDetail `json:"details"`
}

type GetProductsByGenderResponseDetail struct {
	ID              string `json:"productId"`
	Name            string `json:"productName"`
	ImageURL        string `json:"productImage"`
	Price           string `json:"productPrice"`
	AverageRating   string `json:"averageRating"`
	ProductQuantity int    `json:"productQuantity"`
	TotalRating     int    `json:"totalRating"`
}

type IGetProductsByGenderHandler interface {
	Handle(ctx context.Context, req *GetProductsByGenderRequest) (res *GetProductsByGenderResponse, err error)
}
