package domain

import (
	"context"
	"time"
)

type OrderEntity struct {
	OrderID        string    `db:"ORDER_ID" json:"order_id"`
	ProductID      string    `db:"PRODUCT_ID" json:"product_id"`
	UserID         string    `db:"USER_ID" json:"user_id"`
	TransactionID  string    `db:"TRANSACTION_ID" json:"transaction_id"`
	DeliveryStatus string    `db:"DELIVERY_STATUS" json:"delivery_status"`
	PaymentStatus  string    `db:"PAYMENT_STATUS" json:"payment_status"`
	CreatedAt      time.Time `db:"CREATED_AT" json:"created_at"`
	UpdatedAt      time.Time `db:"UPDATED_AT" json:"updated_at"`
}

type IOrderRepo interface {
	Insert(ctx context.Context, orderEntity OrderEntity) error
	Update(ctx context.Context, orderEntity OrderEntity) error
	GetCheckoutItem(ctx context.Context, params GetCheckoutItemParamIn) (GetCheckoutItemResponse, error)
	GetHistory(ctx context.Context, params GetHistoryParamIn) (GetHistoryResponse, error)
}

type UpdateParamIn struct {
	Order OrderEntity
}

type InsertParamIn struct {
	Order OrderEntity
}

type GetHistoryParamIn struct {
	UserID string `json:"user_id"`
}

type GetCheckoutItemParamIn struct {
	UserID string `json:"user_id"`
}
