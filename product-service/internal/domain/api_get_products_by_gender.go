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
	Catergory       string `json:"productCatergory"`
	Price           string `json:"productPrice"`
	Description     string `json:"productDescription"`
	Image           string `json:"productImage"`
	ProductQuantity int    `json:"productQuantity"`
	Gender string `json:"gender"`
}

type IGetProductsByGenderHandler interface {
	Handle(ctx context.Context, req *GetProductsByGenderRequest) (res *GetProductsByGenderResponse, err error)
}
