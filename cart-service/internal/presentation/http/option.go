package http_server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingstonduy/cart-service/internal/domain"
	_ "github.com/kingstonduy/cart-service/resources/docs"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

func (s *HttpServer) WithRoutingOption() option {
	return func(s *HttpServer) error {
		s.App.Post("/add", s.AddCartItem)
		s.App.Post("/get-items", s.GetCart)
		s.App.Post("/update", s.UpdateCartItem)
		s.App.Post("/delete-cart-items", s.DeleteCartItems)

		return nil
	}
}

//	 	@Tags 			CART SERVICE
//		@Summary		GetCart
//		@Description	select * from cart where userID = :1
//		@ID				GetCart
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetCartRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetCartResponse]			"ok"
//		@Router			/is/v1/cart-service/get-items [post]
func (s *HttpServer) GetCart(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetCartRequest, *domain.GetCartResponse](ctx, fiberx.WithAuthentication())
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
func (s *HttpServer) UpdateCartItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.UpdateCartItemRequest, *domain.UpdateCartItemResponse](ctx, fiberx.WithAuthentication())
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
func (s *HttpServer) DeleteCartItems(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.DeleteCartItemsRequest, *domain.DeleteCartItemsResponse](ctx, fiberx.WithAuthentication())
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
func (s *HttpServer) AddCartItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.AddCartItemRequest, *domain.AddCartItemResponse](ctx, fiberx.WithAuthentication())
}
