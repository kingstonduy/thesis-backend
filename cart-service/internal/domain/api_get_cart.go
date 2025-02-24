package domain

import "context"

type GetCartRequest struct {
	UserID string `json:"userId"`
}

type GetCartResponse struct {
	CartItems []GetCartItemDetail `json:"cartItems"`
}

type GetCartItemDetail struct {
	CartItemID       string `json:"cartItemId" bson:"CART_ITEM_ID"`
	ProductID        string `json:"productId" bson:"PRODUCT_ID"`
	ProductName      string `json:"productName" bson:"PRODUCT_NAME"`
	ProductImage     string `json:"productImage" bson:"PRODUCT_IMAGE"`
	ProductCatergory string `json:"productCategory" bson:"PRODUCT_CATEGORY"`
	ProductPrice     string `json:"productPrice" bson:"PRODUCT_PRICE"`
	CartItemQuantity int    `json:"cartItemQuantity" bson:"CART_ITEM_QUANTITY"`
}

type IGetCartHandler interface {
	Handle(ctx context.Context, req *GetCartRequest) (res *GetCartResponse, err error)
}
