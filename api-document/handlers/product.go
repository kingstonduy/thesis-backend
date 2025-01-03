package handlers

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/kingstonduy/api-document/domain/product"
	_ "github.com/kingstonduy/api-document/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

//	 	@Tags 			PRODUCT
//		@Summary		GetAllProduct
//		@Description	list all the products in the inventory
//		@ID				GetAllProduct
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetAllProductRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetAllProductResponse]			"ok"
//		@Router			/is/v1/product-service/get-products [post]
func ListProducts(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetAllProductRequest, *domain.GetAllProductResponse](ctx, errorx.ErrorMessageTimeout)
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
func GetProductDetails(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductDetailRequest, *domain.GetProductDetailResponse](ctx, errorx.ErrorMessageTimeout)
}

//	 	@Tags 			PRODUCT
//		@Summary		ProductExecuteTransaction
//		@Description	remove the product quantity from the inventory
//		@ID				ProductExecuteTransaction
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.ExecuteTransactionRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.ExecuteTransactionResponse]			"ok"
//		@Router			/is/v1/product-service/execute-transaction [post]
func ProductExecuteTransaction(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.ExecuteTransactionRequest, *domain.ExecuteTransactionResponse](ctx, errorx.ErrorMessageTimeout)
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
func RevertTransaction(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.RevertTransactionRequest, *domain.RevertTransactionResponse](ctx, errorx.ErrorMessageTimeout)
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
func GetProductByGender(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductsByGenderRequest, *domain.GetProductsByGenderResponse](ctx, errorx.ErrorMessageTimeout)
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
func GetProductByCategory(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductsByCategoryRequest, *domain.GetProductsByCategoryResponse](ctx, errorx.ErrorMessageTimeout)
}
