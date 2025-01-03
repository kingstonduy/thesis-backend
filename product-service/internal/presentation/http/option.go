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
		s.App.Post("/get-products-by-category", s.GetProductByCategory)
		s.App.Post("/get-products-by-gender", s.GetProductByGender)
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
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetAllProductRequest, *domain.GetAllProductResponse](ctx, fiberx.WithAuthentication())
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
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductDetailRequest, *domain.GetProductDetailResponse](ctx, fiberx.WithAuthentication())
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

//	 	@Tags 			PRODUCT
//		@Summary		GetProductByGender
//		@Description	Get all products filter = {gender}
//		@ID				GetProductByGender
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetProductsByGenderRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetProductsByGenderRequest]			"ok"
//		@Router			/is/v1/product-service/get-products-by-gender [post]
func (s *HttpServer) GetProductByGender(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductsByGenderRequest, *domain.GetProductsByGenderResponse](ctx, fiberx.WithAuthentication())
}

//	 	@Tags 			PRODUCT
//		@Summary		GetProductByCategory
//		@Description	Get all products filter = {category}
//		@ID				GetProductByCategory
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetProductsByCategoryRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetProductsByCategoryResponse]			"ok"
//		@Router			/is/v1/product-service/get-products-by-category [post]
func (s *HttpServer) GetProductByCategory(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductsByCategoryRequest, *domain.GetProductsByCategoryResponse](ctx, fiberx.WithAuthentication())
}
