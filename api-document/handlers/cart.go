package handlers

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/kingstonduy/api-document/domain/cart"
	_ "github.com/kingstonduy/api-document/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

//	 	@Tags 			CART SERVICE
//		@Summary		GetCart
//		@Description	select * from cart where userID = :1
//		@ID				GetCart
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetCartRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetCartResponse]			"ok"
//		@Router			/is/v1/cart-service/get-items [post]
func GetCart(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetCartRequest, *domain.GetCartResponse](ctx, errorx.ErrorMessageTimeout)
}

//	 	@Tags 			CART SERVICE
//		@Summary		UpdateCartItem
//		@Description	Update the quantity of an item in the cart. front end receives ok returns then update StateQuantity
//		@ID				UpdateCartItem
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.UpdateCartItemRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.UpdateCartItemResponse]			"ok"
//		@Router			/is/v1/cart-service/update [post]
func UpdateCartItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.UpdateCartItemRequest, *domain.UpdateCartItemResponse](ctx, errorx.ErrorMessageTimeout)
}

//	 	@Tags 			CART SERVICE
//		@Summary		DeleteCartItem
//		@Description	giving the cartITemID delete the cartItem
//		@ID				DeleteCartItem
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.DeleteCartItemsRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.DeleteCartItemsResponse]			"ok"
//		@Router			/is/v1/cart-service/delete-cart-items [post]
func DeleteCartItems(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.DeleteCartItemsRequest, *domain.DeleteCartItemsResponse](ctx, errorx.ErrorMessageTimeout)
}

//	 	@Tags 			CART SERVICE
//		@Summary		AddCartItem
//		@Description	Insert into cart values ...
//		@ID				AddCartItem
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.AddCartItemRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.AddCartItemResponse]			"ok"
//		@Router			/is/v1/cart-service/add [post]
func AddCartItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.AddCartItemRequest, *domain.AddCartItemResponse](ctx, errorx.ErrorMessageTimeout)
}
