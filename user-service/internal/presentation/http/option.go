package http_server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/transport"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
	"github.com/kingstonduy/user-service/internal/domain"
	_ "github.com/kingstonduy/user-service/resources/docs"
)

func (s *HttpServer) WithRoutingOption() option {
	return func(s *HttpServer) error {
		s.App.Post("/get-user-information", s.GetUserInformation)
		s.App.Post("/update", s.UpdateUserInformation)
		s.App.Post("/register", s.Register)
		s.App.Post("/login", s.Login)
		s.App.Post("/check-jwt", s.CheckJwt)

		return nil
	}
}

//	 	@Tags 			USER SERVICE
//		@Summary		GetUserInformation
//		@Description	Get the user information
//		@ID				GetUserInformation
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetUserInformationRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetUserInformationResponse]			"ok"
//		@Router			/is/v1/user-service/update [post]
func (s *HttpServer) GetUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetUserInformationRequest, *domain.GetUserInformationResponse](ctx, fiberx.WithAuthentication())
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
func (s *HttpServer) UpdateUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.UpdateUserInformationRequest, *domain.UpdateUserInformationResponse](ctx, fiberx.WithAuthentication())
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
func (s *HttpServer) Register(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.RegisterRequest, *domain.RegisterResponse](ctx)
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
func (s *HttpServer) Login(ctx *fiber.Ctx) error {
	f := func(ctx *fiber.Ctx, b []byte) {
		var resType transport.Response[domain.LoginResponse]
		json.Unmarshal(b, &resType)

		//TODO remove EXPIRATION
		jwt, _ := fiberx.CreateToken(resType.Data.UserID)
		ctx.Cookie(&fiber.Cookie{
			Name:    "jwt",
			Value:   jwt,
			Expires: time.Now().Add(time.Hour * 5),
		})

		ctx.Response().Header.Set("jwt", jwt)
	}
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.LoginRequest, *domain.LoginResponse](ctx, fiberx.WithPostHandlerFunc(f))
}

type Req struct {
	Token string `json:"jwt"`
}

type CheckJwtResponse struct {
	UserID string `json:"userId"`
}

func (s *HttpServer) CheckJwt(ctx *fiber.Ctx) error {

	var reqType transport.Request[Req]
	err := ctx.BodyParser(&reqType)
	if err != nil {
		errx := errorx.AuthenticationError("")
		resType := transport.Response[any]{
			Result: transport.Result{
				Code:       errx.Code,
				StatusCode: errx.Status,
				Message:    errx.Message,
				Details:    errx.Details,
			},
		}
		return ctx.Status(errx.Status).JSON(resType)
	}
	token := reqType.Data.Token
	if token == "" {
		errx := errorx.AuthenticationErrorWithDetails("Missing token", "")
		resType := transport.Response[any]{
			Result: transport.Result{
				Code:       errx.Code,
				StatusCode: errx.Status,
				Message:    errx.Message,
				Details:    errx.Details,
			},
		}
		return ctx.Status(errx.Status).JSON(resType)
	}

	userID, err := fiberx.VerifyToken(token)
	if err != nil {
		errx := errorx.AuthenticationErrorWithDetails(err.Error(), "")
		resType := transport.Response[any]{
			Result: transport.Result{
				Code:       errx.Code,
				StatusCode: errx.Status,
				Message:    errx.Message,
				Details:    errx.Details,
			},
		}
		return ctx.Status(errx.Status).JSON(resType)
	}

	resType := transport.Response[any]{
		Result: transport.Result{
			Code:       "00",
			StatusCode: http.StatusOK,
			Message:    "success",
		},
		Data: CheckJwtResponse{
			UserID: userID,
		},
	}
	return ctx.Status(http.StatusOK).JSON(resType)
}
