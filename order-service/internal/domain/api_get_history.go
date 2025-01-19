package domain

import (
	"context"
	"time"
)

type GetHistoryRequest struct {
	UserID string `json:"userID"`
}

type GetHistoryResponse struct {
	Details []GetHistoryResponseDetail `json:"details`
}

type GetHistoryResponseDetail struct {
	ProductID      string    `json:"productId"`
	ProductImage   string    `json:"productImage"`
	ProductName    string    `json:"productName"`
	OrderID        string    `json:"orderId"`
	DeliveryStatus string    `json:"deliveryStatus"`
	PaymentStatus  string    `json:"paymentStatus"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type IGetHistoryHandler interface {
	Handle(ctx context.Context, req *GetHistoryRequest) (res *GetHistoryResponse, err error)
}
