package handlers

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/kingstonduy/api-document/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
	"github.com/shopspring/decimal"
)

type ListProductRequest struct{}
type ListProductResponse struct {
	Products []Product `json:"products"`
}
type Product struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	ImageURL      string          `json:"imageUrl"`
	Price         decimal.Decimal `json:"price"`
	AverageRating decimal.Decimal `json:"averageRating"`
}

//	 	@Tags 			PRODUCT SERVICE
//		@Summary		GetAllProduct
//		@Description	list all the products in the inventory
//		@ID				GetAllProduct
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[ListProductRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[ListProductResponse]			"ok"
//		@Router			/is/v1/product-service/get-products [post]
func ListProducts(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*ListProductRequest, *ListProductResponse](ctx, errorx.ErrorCodeTimeout)
}

type GetProductDetailRequest struct {
	ID string `json:"productId"`
}
type GetProductDetailResponse struct {
	ID          string          `json:"productId"`
	Name        string          `json:"productName"`
	Catergory   string          `json:"productCatergory"`
	Price       decimal.Decimal `json:"productPrice"`
	Description string          `json:"productDescription"`
	Image       string          `json:"productImage"`
}

//	 	@Tags 			PRODUCT SERVICE
//		@Summary		GetProductDetail
//		@Description	Get the detail of a product
//		@ID				GetProductDetail
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[GetProductDetailRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[GetProductDetailResponse]			"ok"
//		@Router			/is/v1/product-service/get-product-detail [post]
func GetProductDetails(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetProductDetailRequest, GetProductDetailResponse](ctx, errorx.ErrorCodeTimeout)
}
