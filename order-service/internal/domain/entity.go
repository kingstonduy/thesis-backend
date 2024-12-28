package domain

import (
	"context"
	"time"
)

type OrderEntity struct {
	OrderID        string    `db:"ORDER_ID" json:"order_id"`
	ProductID      string    `db:"PRODUCT_ID" json:"product_id"`
	UserID         string    `db:"USER_ID" json:"user_id"`
	OrderStatus    string    `db:"ORDER_STATUS" json:"order_status"`
	DeliveryStatus string    `db:"DELIVERY_STATUS" json:"delivery_status"`
	PaymentStatus  string    `db:"PAYMENT_STATUS" json:"payment_status"`
	CreatedAt      time.Time `db:"CREATED_AT" json:"created_at"`
	UpdatedAt      time.Time `db:"UPDATED_AT" json:"updated_at"`
}

type IOrderRepo interface {
	Insert(ctx context.Context, params InsertParamIn) error
	Update(ctx context.Context, params UpdateParamIn) error
	GetCheckoutItem(ctx context.Context, params GetCheckoutItemParamIn) (GetCheckoutItemParamOut, error)
	GetHistory(ctx context.Context, params GetHistoryParamIn) (GetHistoryParamOut, error)
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

type GetHistoryParamOut struct {
	Order OrderEntity
}

type GetCheckoutItemParamIn struct {
	UserID string `json:"user_id"`
}

type GetCheckoutItemParamOut struct{}
