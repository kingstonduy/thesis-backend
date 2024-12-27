package handlers

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/kingstonduy/api-document/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
	"github.com/shopspring/decimal"
)

type GetCheckoutItemRequest struct {
	UserID string `json:"userID"`
}
type GetCheckoutItemResponse struct {
	CheckoutItemDetails []GetCheckoutItemDetail
}

type GetCheckoutItemDetail struct {
	ProductID        string          `json:"productId"`
	ProductImage     string          `json:"productImage"`
	ProductName      string          `json:"productName"`
	ProductCatergory string          `json:"productCatergory"`
	ProductQuantity  int             `json:"productQuantity"`
	PricePerUnit     decimal.Decimal `json:"pricePerUnit"`
}

//	 	@Tags 			ORDER SERVICE
//		@Summary		GetCheckoutItem
//		@Description	Get the checkout item (cdc from cart table)
//		@ID				GetCheckoutItem
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[GetCheckoutItemRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[GetCheckoutItemResponse]			"ok"
//		@Router			/is/v1/order-service/get-checkout-item [post]
func GetCheckoutItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetCheckoutItemRequest, *GetCheckoutItemResponse](ctx, errorx.ErrorCodeTimeout)
}

type CheckoutRequest struct {
	UserID          string           `json:"userID"`
	CheckoutDetails []CheckoutDetail `json:"checkoutDetails"`
}
type CheckoutResponse struct {
}

type CheckoutDetail struct {
	ProductID       string `json:"productId"`
	ProductQuantity int    `json:"quantity"`
}

//	 	@Tags 			ORDER SERVICE
//		@Summary		Checkout
//		@Description	Delete all items on cart -> minus quantity of product in cart service
//		@ID				Checkout
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[CheckoutRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[CheckoutResponse]			"ok"
//		@Router			/is/v1/order-service/checkout [post]
func Checkout(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*CheckoutRequest, *CheckoutResponse](ctx, errorx.ErrorCodeTimeout)
}

type GetPurchasedProductsRequest struct {
	UserID string `json:"userID"`
}
type GetPurchasedProductsResponse struct {
	PurchasedProducts []PurchasedProductDetail `json:"purchasedProducts"`
}
type PurchasedProductDetail struct {
	ProductID      string `json:"productId"`
	ProductImage   string `json:"productImage"`
	ProductName    string `json:"productName"`
	OrderID        string `json:"orderId"`
	DeliveryStatus string `json:"deliveryStatus"`
	PaymentStatus  string `json:"paymentStatus"`
}

//	 	@Tags 			ORDER SERVICE
//		@Summary		GetPurchasedProducts
//		@Description	Show a history of purchased products from a user
//		@ID				GetPurchasedProducts
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[GetPurchasedProductsRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[GetPurchasedProductsResponse]			"ok"
//		@Router			/is/v1/order-service/get-history [post]
func GetPurchasedProducts(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetPurchasedProductsRequest, *GetPurchasedProductsResponse](ctx, errorx.ErrorCodeTimeout)
}
