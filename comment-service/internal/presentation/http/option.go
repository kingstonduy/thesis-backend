package http_server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingstonduy/comment-service/internal/domain"
	_ "github.com/kingstonduy/comment-service/resources/docs"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
)

func (s *HttpServer) WithRoutingOption() option {
	return func(s *HttpServer) error {
		s.App.Post("/get-comment", s.GetComment)
		s.App.Post("/add", s.AddComment)

		return nil
	}
}

//	 	@Tags 			COMMENT SERVICE
//		@Summary		GetComment
//		@Description	select * from table comment where productID = :1
//		@ID				GetComment
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetCommentRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetCommentResponse]			"ok"
//		@Router			/is/v1/comment-service/get-comment [post]
func (s *HttpServer) GetComment(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetCommentRequest, *domain.GetCommentResponse](ctx, fiberx.WithAuthentication())
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
func (s *HttpServer) AddComment(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.AddCommentRequest, *domain.AddCommentResponse](ctx, fiberx.WithAuthentication())
}
