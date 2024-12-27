package handlers

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/kingstonduy/api-document/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

type GetCartRequest struct {
	UserID string `json:"userId"`
}
type GetCartResponse struct {
	CartItems []CartItem `json:"cartItems"`
}
type CartItem struct {
	CartItemID       string `json:"cartItemId"`
	ProductID        string `json:"productId"`
	ProductName      string `json:"productName"`
	ProductImage     string `json:"productImage"`
	ProductCatergory string `json:"productCatergory"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

//	 	@Tags 			CART SERVICE
//		@Summary		GetCart
//		@Description	select * from cart where userID = :1
//		@ID				GetCart
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[GetCartRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[GetCartResponse]			"ok"
//		@Router			/is/v1/cart-service/get-items [post]
func GetCart(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetCartRequest, *GetCartResponse](ctx, errorx.ErrorCodeTimeout)
}

type UpdateCartItemRequest struct {
	CartItemID       string `json:"cartItemId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type UpdateCartItemResponse struct {
}

//	 	@Tags 			CART SERVICE
//		@Summary		UpdateCartItem
//		@Description	Update the quantity of an item in the cart. front end receives ok returns then update StateQuantity
//		@ID				UpdateCartItem
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[UpdateCartItemRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[UpdateCartItemResponse]			"ok"
//		@Router			/is/v1/cart-service/update [post]
func UpdateCartItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*UpdateCartItemRequest, *UpdateCartItemResponse](ctx, errorx.ErrorCodeTimeout)
}

type DeleteCartItemRequest struct {
	CartItems []DeleteCartItemDetail `json:"cartItems"`
}

type DeleteCartItemDetail struct {
	UserID string `json:"userID"`
}

type DeleteCartItemResposne struct{}

//	 	@Tags 			CART SERVICE
//		@Summary		DeleteCartItem
//		@Description	DELETE FROM cart WHERE userID = ?;
//		@ID				DeleteCartItem
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[DeleteCartItemRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[DeleteCartItemResposne]			"ok"
//		@Router			/is/v1/cart-service/delete [post]
func DeleteCartItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*DeleteCartItemRequest, *DeleteCartItemResposne](ctx, errorx.ErrorCodeTimeout)
}

type AddCartItemRequest struct {
	CartItems []AddCartItemDetail `json:"cartItems"`
}

type AddCartItemDetail struct {
	CartItemID       string `json:"cartItemId"`
	ProductID        string `json:"productId"`
	ProductName      string `json:"productName"`
	ProductImage     string `json:"productImage"`
	ProductCatergory string `json:"productCatergory"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}
type AddCartItemResposne struct {
}

//	 	@Tags 			CART SERVICE
//		@Summary		AddCartItem
//		@Description	Insert into cart values ...
//		@ID				AddCartItem
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[AddCartItemRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[AddCartItemResposne]			"ok"
//		@Router			/is/v1/cart-service/add [post]
func AddCartItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*AddCartItemRequest, *AddCartItemResposne](ctx, errorx.ErrorCodeTimeout)
}
