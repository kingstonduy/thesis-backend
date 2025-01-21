package configuration

import (
	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/pipeline"
)

func ResgisterPipeline(
	IAddCartItemHandler domain.IAddCartItemHandler,
	IGetCartHandler domain.IGetCartHandler,
	IUpdateCartItemHandler domain.IUpdateCartItemHandler,
) {
	pipeline.RegisterRequestHandler(IAddCartItemHandler)
	pipeline.RegisterRequestHandler(IGetCartHandler)
	pipeline.RegisterRequestHandler(IUpdateCartItemHandler)
}
