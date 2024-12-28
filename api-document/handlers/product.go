package handlers

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/kingstonduy/api-document/domain/product"
	_ "github.com/kingstonduy/api-document/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

//	 	@Tags 			product
//		@Summary		GetAllProduct
//		@Description	list all the products in the inventory
//		@ID				GetAllProduct
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetAllProductRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetAllProductResponse]			"ok"
//		@Router			/is/v1/product-service/get-products [post]
func GetAllProduct(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetAllProductRequest, *domain.GetAllProductResponse](ctx, errorx.ErrorCodeTimeout)
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
func GetProductDetails(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductDetailRequest, *domain.GetProductDetailResponse](ctx, errorx.ErrorCodeTimeout)
}
