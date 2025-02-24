package domain

import "context"

type CartItemEvent struct {
	CartItemID       string `json:"CART_ITEM_ID"`
	UserID           string `json:"USER_ID"`
	ProductID        string `json:"PRODUCT_ID"`
	CartItemQuantity int    `json:"CART_ITEM_QUANTITY"`
	Deleted          bool   `json:"__deleted,string"`
}

type CartItemEventRes struct {
}

type ICartItemEventHandler interface {
	Handle(ctx context.Context, req Event[*CartItemEvent]) (res *CartItemEventRes, err error)
}
