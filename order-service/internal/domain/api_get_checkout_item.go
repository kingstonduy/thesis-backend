package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

type GetCheckoutItemRequest struct {
	UserID string `json:"userID"`
}

type GetCheckoutItemResponse struct {
	Details []GetCheckoutItemResponseDetail `json:"details"`
}

type GetCheckoutItemResponseDetail struct {
	ProductID        string          `json:"productId"`
	ProductImage     string          `json:"productImage"`
	ProductName      string          `json:"productName"`
	ProductCatergory string          `json:"productCatergory"`
	ProductQuantity  int             `json:"productQuantity"`
	PricePerUnit     decimal.Decimal `json:"pricePerUnit"`
}

type IGetCheckoutItemHandler interface {
	Handle(ctx context.Context, req *GetCheckoutItemRequest) (res *GetCheckoutItemResponse, err error)
}
