package http_server

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
	"github.com/kingstonduy/product-service/internal/domain"
	_ "github.com/kingstonduy/product-service/resources/docs"
)

func (s *HttpServer) WithRoutingOption() option {
	return func(s *HttpServer) error {
		s.App.Post("/get-products", s.ListProducts)
		s.App.Post("/get-product-detail", s.GetProductDetails)
		s.App.Post("/execute-transaction", s.ExecuteTransaction)
		s.App.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("CORS Configured!")
		})
		return nil
	}
}

//	 	@Tags 			PRODUCT
//		@Summary		GetAllProduct
//		@Description	list all the products in the inventory
//		@ID				GetAllProduct
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetAllProductRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetAllProductResponse]			"ok"
//		@Router			/is/v1/product-service/get-products [post]
func (s *HttpServer) ListProducts(ctx *fiber.Ctx) error {
	// return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetAllProductRequest, *domain.GetAllProductResponse](ctx, fiberx.WithAuthentication())
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetAllProductRequest, *domain.GetAllProductResponse](ctx)
}

//	 	@Tags 			PRODUCT
//		@Summary		GetProductDetail
//		@Description	Get the detail of a product
//		@ID				GetProductDetail
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetProductDetailRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetProductDetailResponse]			"ok"
//		@Router			/is/v1/product-service/get-product-detail [post]
func (s *HttpServer) GetProductDetails(ctx *fiber.Ctx) error {
	// return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductDetailRequest, *domain.GetProductDetailResponse](ctx, fiberx.WithAuthentication())
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductDetailRequest, *domain.GetProductDetailResponse](ctx)
}

//	 	@Tags 			PRODUCT
//		@Summary		ExecuteTransaction
//		@Description	remove the product quantity from the inventory
//		@ID				ExecuteTransaction
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.ExecuteTransactionRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.ExecuteTransactionResponse]			"ok"
//		@Router			/is/v1/product-service/execute-transaction [post]
func (s *HttpServer) ExecuteTransaction(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.ExecuteTransactionRequest, *domain.ExecuteTransactionResponse](ctx, fiberx.WithAuthentication())
}

//	 	@Tags 			PRODUCT
//		@Summary		RevertTransaction
//		@Description	revert the transaction
//		@ID				RevertTransaction
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.RevertTransactionRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.RevertTransactionResponse]			"ok"
//		@Router			/is/v1/product-service/revert-transaction [post]
func (s *HttpServer) RevertTransaction(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.RevertTransactionRequest, *domain.RevertTransactionResponse](ctx, fiberx.WithAuthentication())
}
