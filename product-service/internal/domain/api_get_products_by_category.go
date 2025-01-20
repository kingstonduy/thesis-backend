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
	Catergory       string `json:"productCatergory"`
	Price           string `json:"productPrice"`
	Description     string `json:"productDescription"`
	Image           string `json:"productImage"`
	ProductQuantity int    `json:"productQuantity"`
	Gender string `json:"gender"`
}

type IGetProductsByCategoryHandler interface {
	Handle(ctx context.Context, req *GetProductsByCategoryRequest) (res *GetProductsByCategoryResponse, err error)
}
