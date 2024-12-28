package configuration

import (
	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/pipeline"
)

func ResgisterPipeline(
	AddCartItemHandler domain.AddCartItemHandler,
	GetCartHandler domain.GetCartHandler,
	UpdateCartItemHandler domain.UpdateCartItemHandler,
	DeleteCartItemHandler domain.DeleteCartItemHandler,
	DeleteUserCartHandler domain.DeleteUserCartHandler,
) {
	pipeline.RegisterRequestHandler(AddCartItemHandler)
	pipeline.RegisterRequestHandler(DeleteCartItemHandler)
	pipeline.RegisterRequestHandler(GetCartHandler)
	pipeline.RegisterRequestHandler(UpdateCartItemHandler)
	pipeline.RegisterRequestHandler(DeleteUserCartHandler)
}
