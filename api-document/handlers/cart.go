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
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetCartRequest, *domain.GetCartResponse](ctx, errorx.ErrorCodeTimeout)
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
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.UpdateCartItemRequest, *domain.UpdateCartItemResponse](ctx, errorx.ErrorCodeTimeout)
}

//	 	@Tags 			CART SERVICE
//		@Summary		DeleteCartItem
//		@Description	giving the cartITemID delete the cartItem
//		@ID				DeleteCartItem
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.DeleteCartItemRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.DeleteCartItemResponse]			"ok"
//		@Router			/is/v1/cart-service/delete-cart-item [post]
func DeleteCartItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.DeleteCartItemRequest, *domain.DeleteCartItemResponse](ctx, errorx.ErrorCodeTimeout)
}

//	 	@Tags 			CART SERVICE
//		@Summary		DeleteUserCart
//		@Description	delete all the record in table cart with corresponding userID
//		@ID				DeleteUserCart
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.DeleteUserCartRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.DeleteUserCartResponse]			"ok"
//		@Router			/is/v1/cart-service/delete-user-cart [post]
func DeleteUserCart(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.DeleteUserCartRequest, *domain.DeleteUserCartResponse](ctx, errorx.ErrorCodeTimeout)
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
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.AddCartItemRequest, *domain.AddCartItemResponse](ctx, errorx.ErrorCodeTimeout)
}
