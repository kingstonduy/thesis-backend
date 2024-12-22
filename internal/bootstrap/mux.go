package configuration

import (
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/thesis-backend/internal/domain"
)

func ResgisterPipeline(
	IGetProductsHandler domain.IGetProductsHandler,
	IGetProductDetailHandler domain.IGetProductDetailHandler,
) {
	pipeline.RegisterRequestHandler(IGetProductsHandler)
	pipeline.RegisterRequestHandler(IGetProductDetailHandler)

}
