package handlers

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/kingstonduy/api-document/domain/user"
	_ "github.com/kingstonduy/api-document/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

//	 	@Tags 			USER SERVICE
//		@Summary		GetUserInformation
//		@Description	Get the user information
//		@ID				GetUserInformation
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetUserInformationRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetUserInformationResponse]			"ok"
//		@Router			/is/v1/user-service/update [post]
func GetUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetUserInformationRequest, *domain.GetUserInformationResponse](ctx, errorx.ErrorCodeTimeout)
}

//	 	@Tags 			USER SERVICE
//		@Summary		UpdateUserInformation
//		@Description	Update the user information
//		@ID				UpdateUserInformation
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.UpdateUserInformationRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.UpdateUserInformationResponse]			"ok"
//		@Router			/is/v1/user-service/get-product-detail [post]
func UpdateUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.UpdateUserInformationRequest, *domain.UpdateUserInformationResponse](ctx, errorx.ErrorCodeTimeout)
}

//	 	@Tags 			USER SERVICE
//		@Summary		RegisterUser
//		@Description	Register a new user
//		@ID				RegisterUser
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.RegisterRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.RegisterResponse]			"ok"
//		@Router			/is/v1/user-service/register [post]
func Register(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.RegisterRequest, *domain.RegisterResponse](ctx, errorx.ErrorCodeTimeout)
}

//	 	@Tags 			USER SERVICE
//		@Summary		Login
//		@Description	login for a new session
//		@ID				Login
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.LoginRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.LoginResponse]			"ok"
//		@Router			/is/v1/user-service/login [post]
func Login(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.LoginRequest, *domain.LoginResponse](ctx, errorx.ErrorCodeTimeout)
}
