package domain

import (
	"context"
)

type GetAllProductPageRequest struct {
	PageNumber int `json:"pageNumber"`
}

type GetAllProductPageResponse struct {
	TotalPage int                               `json:"totalPage"`
	Details   []GetAllProductPageResponseDetail `json:"details"`
}

type GetAllProductPageResponseDetail struct {
	ID              string `json:"productId"`
	Name            string `json:"productName"`
	Catergory       string `json:"productCatergory"`
	Price           string `json:"productPrice"`
	Description     string `json:"productDescription"`
	Image           string `json:"productImage"`
	ProductQuantity int    `json:"productQuantity"`
	Gender          string `json:"gender"`
}
type IGetProductsPageHandler interface {
	Handle(ctx context.Context, req *GetAllProductPageRequest) (res *GetAllProductPageResponse, err error)
}
