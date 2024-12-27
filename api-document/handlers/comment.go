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
//		@Description	select * from table comment where productID = :1
//		@ID				GetCommentsByProductID
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[GetCommentsByProductIDRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[GetCommentsByProductIDResponse]			"ok"
//		@Router			/is/v1/comment-service/get-comment [post]
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
//		@Description	insert into table comment
//		@ID				AddComment
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[AddCommentRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[AddCommentResponse]			"ok"
//		@Router			/is/v1/comment-service/add [post]
func AddComment(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*AddCommentRequest, *AddCommentResponse](ctx, errorx.ErrorCodeTimeout)
}
