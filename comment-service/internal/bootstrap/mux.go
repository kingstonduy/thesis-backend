package configuration

import (
	"github.com/kingstonduy/comment-service/internal/domain"
	"github.com/kingstonduy/go-core/pipeline"
)

func RegisterPipeline(
	IAddCommentHandler domain.IAddCommentHandler,
	IGetCommentHandler domain.IGetCommentHandler,
) {
	pipeline.RegisterRequestHandler(IAddCommentHandler)
	pipeline.RegisterRequestHandler(IGetCommentHandler)
}
