package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

type GetProductDetailRequest struct {
	ID string `json:"productId"`
}

type GetProductDetailResponse struct {
	ID              string          `json:"productId"`
	Name            string          `json:"productName"`
	Catergory       string          `json:"productCatergory"`
	Price           decimal.Decimal `json:"productPrice"`
	Description     string          `json:"productDescription"`
	Image           string          `json:"productImage"`
	ProductQuantity int             `json:"productQuantity"`
}

type IGetProductDetailHandler interface {
	Handle(ctx context.Context, req *GetProductDetailRequest) (res *GetProductDetailResponse, err error)
}
