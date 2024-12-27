package http_server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingstonduy/cart-service/internal/domain"
	_ "github.com/kingstonduy/cart-service/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

func (s *HttpServer) WithRoutingOption() option {
	return func(s *HttpServer) error {
		s.App.Post("/get-products", s.ListProducts)
		s.App.Post("/get-product-detail", s.GetProductDetails)
		return nil
	}
}

//	 	@Tags 			product
//		@Summary		GetAllProduct
//		@Description	list all the products in the inventory
//		@ID				GetAllProduct
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetProductsRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetProductsResponse]			"ok"
//		@Router			/is/v1/product-service/get-products [post]
func (s *HttpServer) ListProducts(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductsRequest, *domain.GetProductsResponse](ctx, errorx.ErrorCodeTimeout)
}

//	 	@Tags 			product
//		@Summary		GetProductDetail
//		@Description	Get the detail of a product
//		@ID				GetProductDetail
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetProductDetailRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetProductDetailResponse]			"ok"
//		@Router			/is/v1/product-service/get-product-detail [post]
func (s *HttpServer) GetProductDetails(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductDetailRequest, *domain.GetProductDetailResponse](ctx, errorx.ErrorCodeTimeout)
}
