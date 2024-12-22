package configuration

import (
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/product-service/internal/domain"
)

func ResgisterPipeline(
	IGetProductsHandler domain.IGetProductsHandler,
	IGetProductDetailHandler domain.IGetProductDetailHandler,
) {
	pipeline.RegisterRequestHandler(IGetProductsHandler)
	pipeline.RegisterRequestHandler(IGetProductDetailHandler)

}
