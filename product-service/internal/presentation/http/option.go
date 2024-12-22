package http_server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
	"github.com/kingstonduy/product-service/internal/domain"
	_ "github.com/kingstonduy/product-service/resources/docs"
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
//		@Router			/is/v1/product/get-products [post]
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
//		@Router			/is/v1/product/get-product-detail [post]
func (s *HttpServer) GetProductDetails(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetProductDetailRequest, *domain.GetProductDetailResponse](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           comment
// @Summary        GetComments
// @Description    Get comments for a product
// @ID             GetComments
// @Produce        json
// @Param          productId path string true "Product ID"
// @Success        200 {object} transport.Response[[]domain.Comment] "ok"
// @Router         /products/{productId}/comments [get]
func (s *HttpServer) GetComments(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[any, []domain.Comment](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           comment
// @Summary        AddComment
// @Description    Add a comment to a product
// @ID             AddComment
// @Accept         json
// @Produce        json
// @Param          request body transport.Request[domain.AddCommentRequest] true "Request"
// @Success        200 {object} transport.Response[domain.AddCommentResponse] "ok"
// @Router         /products/{productId}/comments [post]
func (s *HttpServer) AddComment(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.AddCommentRequest, *domain.AddCommentResponse](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           cart
// @Summary        ListCartItems
// @Description    Get all items in the cart
// @ID             ListCartItems
// @Produce        json
// @Success        200 {object} transport.Response[[]domain.CartItem] "ok"
// @Router         /cart [get]
func (s *HttpServer) ListCartItems(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[any, []domain.CartItem](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           cart
// @Summary        UpdateCartItemQuantity
// @Description    Update quantity of a product in the cart
// @ID             UpdateCartItemQuantity
// @Accept         json
// @Produce        json
// @Param          request body transport.Request[domain.UpdateCartRequest] true "Request"
// @Success        200 {object} transport.Response[domain.UpdateCartResponse] "ok"
// @Router         /cart/{productId} [put]
func (s *HttpServer) UpdateCartItemQuantity(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.UpdateCartRequest, *domain.UpdateCartResponse](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           cart
// @Summary        Checkout
// @Description    Checkout cart items
// @ID             Checkout
// @Accept         json
// @Produce        json
// @Success        200 {object} transport.Response[domain.CheckoutResponse] "ok"
// @Router         /cart/checkout [post]
func (s *HttpServer) Checkout(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[any, *domain.CheckoutResponse](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           order
// @Summary        ListPurchasedProducts
// @Description    Get all purchased products
// @ID             ListPurchasedProducts
// @Produce        json
// @Success        200 {object} transport.Response[[]domain.PurchasedProduct] "ok"
// @Router         /orders [get]
func (s *HttpServer) ListPurchasedProducts(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[any, []domain.PurchasedProduct](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           user
// @Summary        GetUserInformation
// @Description    Get user details
// @ID             GetUserInformation
// @Produce        json
// @Success        200 {object} transport.Response[domain.UserInformation] "ok"
// @Router         /users/{userId} [get]
func (s *HttpServer) GetUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[any, *domain.UserInformation](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           user
// @Summary        UpdateUserInformation
// @Description    Update user details
// @ID             UpdateUserInformation
// @Accept         json
// @Produce        json
// @Param          request body transport.Request[domain.UpdateUserRequest] true "Request"
// @Success        200 {object} transport.Response[domain.UpdateUserResponse] "ok"
// @Router         /users/{userId} [put]
func (s *HttpServer) UpdateUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.UpdateUserRequest, *domain.UpdateUserResponse](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           auth
// @Summary        Login
// @Description    Authenticate user
// @ID             Login
// @Accept         json
// @Produce        json
// @Param          request body transport.Request[domain.LoginRequest] true "Request"
// @Success        200 {object} transport.Response[domain.LoginResponse] "ok"
// @Router         /auth/login [post]
func (s *HttpServer) Login(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.LoginRequest, *domain.LoginResponse](ctx, errorx.ErrorCodeTimeout)
}

// @Tags           auth
// @Summary        ValidateJWT
// @Description    Validate JWT token
// @ID             ValidateJWT
// @Accept         json
// @Produce        json
// @Param          request body transport.Request[domain.ValidateTokenRequest] true "Request"
// @Success        200 {object} transport.Response[domain.ValidateTokenResponse] "ok"
// @Router         /auth/validate [post]
func (s *HttpServer) ValidateJWT(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.ValidateTokenRequest, *domain.ValidateTokenResponse](ctx, errorx.ErrorCodeTimeout)
}
