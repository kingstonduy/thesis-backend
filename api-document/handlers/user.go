package handlers

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/kingstonduy/api-document/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

type GetUserInformationRequest struct {
	UserID string `json:"userId"`
}
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
//		@Router			/is/v1/user-service/get-user-information [post]
func GetUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*GetUserInformationRequest, *GetUserInformationResponse](ctx, errorx.ErrorCodeTimeout)
}

type UpdateUserInformationRequest struct {
	UserID       string `json:"userId"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Gender       string `json:"gender"`
	DateOfBirth  string `json:"dateOfBirth"`
	Street       string `json:"street"`
	City         string `json:"city"`
	CityCode     string `json:"cityCode"`
	District     string `json:"district"`
	DistrictCode string `json:"districtCode"`
	Ward         string `json:"ward"`
	WardCode     string `json:"wardCode"`
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
//		@Router			/is/v1/user-service/update [post]
func UpdateUserInformation(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*UpdateUserInformationRequest, *UpdateUserInformationResponse](ctx, errorx.ErrorCodeTimeout)
}

type RegisterUserRequest struct {
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Gender       string `json:"gender"`
	DateOfBirth  string `json:"dateOfBirth"`
	Street       string `json:"street"`
	City         string `json:"city"`
	CityCode     string `json:"cityCode"`
	District     string `json:"district"`
	DistrictCode string `json:"districtCode"`
	Ward         string `json:"ward"`
	WardCode     string `json:"wardCode"`
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
//		@Router			/is/v1/user-service/register [post]
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
//		@Router			/is/v1/user-service/login [post]
func Login(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*LoginRequest, *LoginResponse](ctx, errorx.ErrorCodeTimeout)
}
