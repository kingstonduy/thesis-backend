package handlers

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/kingstonduy/api-document/domain/comment"
	_ "github.com/kingstonduy/api-document/resources/docs"
	"github.com/kingstonduy/go-core/errorx"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

//	 	@Tags 			COMMENT SERVICE
//		@Summary		GetComment
//		@Description	select * from table comment where productID = :1
//		@ID				GetComment
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetCommentRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetCommentResponse]			"ok"
//		@Router			/is/v1/comment-service/get-comment [post]
func GetComment(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetCommentRequest, *domain.GetCommentResponse](ctx, errorx.ErrorCodeTimeout)
}

//	 	@Tags 			COMMENT SERVICE
//		@Summary		AddComment
//		@Description	insert into table comment
//		@ID				AddComment
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.AddCommentRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.AddCommentResponse]			"ok"
//		@Router			/is/v1/comment-service/add [post]
func AddComment(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.AddCommentRequest, *domain.AddCommentResponse](ctx, errorx.ErrorCodeTimeout)
}
