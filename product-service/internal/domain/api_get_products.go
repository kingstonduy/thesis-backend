package domain

import (
	"context"
)

type GetAllProductRequest struct {
}

type GetAllProductResponse struct {
	Details []GetAllProductResponseDetail `json:"details"`
}

type GetAllProductResponseDetail struct {
	ID              string `json:"productId"`
	Name            string `json:"productName"`
	Catergory       string `json:"productCatergory"`
	Price           string `json:"productPrice"`
	Description     string `json:"productDescription"`
	Image           string `json:"productImage"`
	ProductQuantity int    `json:"productQuantity"`
	Gender string `json:"gender"`
}
type IGetProductsHandler interface {
	Handle(ctx context.Context, req *GetAllProductRequest) (res *GetAllProductResponse, err error)
}
