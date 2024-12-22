package handlers

import (
	"time"

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
//		@Router			/is/v1/product/get-products [post]
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
//		@Router			/is/v1/product/get-product-detail [post]
func GetProductDetails(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetProductDetailRequest, GetProductDetailResponse](ctx, errorx.ErrorCodeTimeout)
}

type GetCommentsByProductIDRequest struct {
	ProductID string `json:"productId"`
}
type GetCommentsByProductIDResponse struct {
	Comments []Comment `json:"comments"`
}
type Comment struct {
	UserImage string          `json:"userImage"`
	Username  string          `json:"userName"`
	Timestamp time.Time       `json:"timestamp"`
	Content   string          `json:"content"`
	Rating    decimal.Decimal `json:"rating"`
}

//	 	@Tags 			COMMENT SERVICE
//		@Summary		GetCommentsByProductID
//		@Description	Get all the comments related to a product
//		@ID				GetCommentsByProductID
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[GetCommentsByProductIDRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[GetCommentsByProductIDResponse]			"ok"
//		@Router			/is/v1/comment/get-product-detail [post]
func GetCommentsByProductID(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetCommentsByProductIDRequest, *GetCommentsByProductIDResponse](ctx, errorx.ErrorCodeTimeout)
}

type AddCommentRequest struct {
	ProductID string          `json:"productId"`
	UserID    string          `json:"userId"`
	Content   string          `json:"content"`
	Rating    decimal.Decimal `json:"rating"`
}
type AddCommentResponse struct{}

//	 	@Tags 			COMMENT SERVICE
//		@Summary		AddComment
//		@Description	Add comment description
//		@ID				AddComment
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[AddCommentRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[AddCommentResponse]			"ok"
//		@Router			/is/v1/comment/add-comment [post]
func AddComment(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*AddCommentRequest, *AddCommentResponse](ctx, errorx.ErrorCodeTimeout)
}

type GetCartRequest struct{}
type GetCartResponse struct {
	CartItems []CartItem `json:"cartItems"`
}
type CartItem struct {
	CartItemID       string `json:"cartItemId"`
	ProductID        string `json:"productId"`
	ProductName      string `json:"productName"`
	ProductImage     string `json:"productImage"`
	ProductCatergory string `json:"productCatergory"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

//	 	@Tags 			CART SERVICE
//		@Summary		GetCart
//		@Description	Get all items on user's cart
//		@ID				GetCart
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[GetCartRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[GetCartResponse]			"ok"
//		@Router			/is/v1/cart/get-items [post]
func GetCart(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetCartRequest, *GetCartResponse](ctx, errorx.ErrorCodeTimeout)
}

type UpdateCartItemRequest struct {
	CartItemID       string `json:"cartItemId"`
	CartItemQuantity int    `json:"cartItemQuantity"`
}

type UpdateCartItemResponse struct {
}

//	 	@Tags 			CART SERVICE
//		@Summary		UpdateCartItem
//		@Description	Update the quantity of an item in the cart
//		@ID				UpdateCartItem
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[UpdateCartItemRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[UpdateCartItemResponse]			"ok"
//		@Router			/is/v1/cart/update [post]
func UpdateCartItem(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*UpdateCartItemRequest, *UpdateCartItemResponse](ctx, errorx.ErrorCodeTimeout)
}

type CheckoutRequest struct {
	CheckoutItems []CheckoutItem `json:"checkoutItems"`
}
type CheckoutResponse struct{}

type CheckoutItem struct {
	CartItemID string `json:"cartItemId"`
}

//	 	@Tags 			ORDER SERVICE
//		@Summary		Checkout
//		@Description	Checkout the cart
//		@ID				Checkout
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[CheckoutRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[CheckoutResponse]			"ok"
//		@Router			/is/v1/order/checkout [post]
func Checkout(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*CheckoutRequest, *CheckoutResponse](ctx, errorx.ErrorCodeTimeout)
}

type GetPurchasedProductsRequest struct{}
type GetPurchasedProductsResponse struct {
	ProductID      string `json:"productId"`
	ProductImage   string `json:"productImage"`
	ProductName    string `json:"productName"`
	OrderID        string `json:"orderId"`
	DeliveryStatus string `json:"deliveryStatus"`
	PaymentStatus  string `json:"paymentStatus"`
}

//	 	@Tags 			ORDER SERVICE
//		@Summary		GetPurchasedProducts
//		@Description	Get all the purchased products
//		@ID				GetPurchasedProducts
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[GetPurchasedProductsRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[GetPurchasedProductsResponse]			"ok"
//		@Router			/is/v1/order/get-history [post]
func GetPurchasedProducts(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetPurchasedProductsRequest, *GetPurchasedProductsResponse](ctx, errorx.ErrorCodeTimeout)
}

type GetUserInformationRequest struct{}
type GetUserInformationResponse struct {
	UserID      string `json:"userId"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"dateOfBirth"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Ward        string `json:"ward"`
}

//	 	@Tags 			USER SERVICE
//		@Summary		GetUserInformation
//		@Description	Get the user information
//		@ID				GetUserInformation
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[GetUserInformationRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[GetUserInformationResponse]			"ok"
//		@Router			/is/v1/user/get-user-information [post]
func GetUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetUserInformationRequest, *GetUserInformationResponse](ctx, errorx.ErrorCodeTimeout)
}

type UpdateUserInformationRequest struct {
	UserID      string `json:"userId"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"dateOfBirth"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Ward        string `json:"ward"`
}
type UpdateUserInformationResponse struct{}

//	 	@Tags 			USER SERVICE
//		@Summary		UpdateUserInformation
//		@Description	Update the user information
//		@ID				UpdateUserInformation
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[UpdateUserInformationRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[UpdateUserInformationResponse]			"ok"
//		@Router			/is/v1/user/get-product-detail [post]
func UpdateUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*UpdateUserInformationRequest, *UpdateUserInformationResponse](ctx, errorx.ErrorCodeTimeout)
}

type RegisterUserRequest struct {
	UserID      string `json:"userId"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"dateOfBirth"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Ward        string `json:"ward"`
}
type RegisterUserResponse struct{}

//	 	@Tags 			USER SERVICE
//		@Summary		RegisterUser
//		@Description	Register a new user
//		@ID				RegisterUser
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[RegisterUserRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[RegisterUserResponse]			"ok"
//		@Router			/is/v1/user/register [post]
func RegisterUser(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*RegisterUserRequest, *RegisterUserResponse](ctx, errorx.ErrorCodeTimeout)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct{}

//	 	@Tags 			USER SERVICE
//		@Summary		Login
//		@Description	login for a new session
//		@ID				Login
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[LoginRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[LoginResponse]			"ok"
//		@Router			/is/v1/user/login [post]
func Login(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*LoginRequest, *LoginResponse](ctx, errorx.ErrorCodeTimeout)
}
