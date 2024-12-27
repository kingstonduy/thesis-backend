package configuration

import (
	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/pipeline"
)

func ResgisterPipeline(
	AddCartItemHandler domain.AddCartItemHandler,
	DeleteCartItemHandler domain.DeleteCartItemHandler,
	GetCartHandler domain.GetCartHandler,
	UpdateCartItemHandler domain.UpdateCartItemHandler,
) {
	pipeline.RegisterRequestHandler(AddCartItemHandler)
	pipeline.RegisterRequestHandler(DeleteCartItemHandler)
	pipeline.RegisterRequestHandler(GetCartHandler)
	pipeline.RegisterRequestHandler(UpdateCartItemHandler)
}
