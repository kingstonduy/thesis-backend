package configuration

import (
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/product-service/internal/domain"
)

func ResgisterPipeline(
	IGetProductsHandler domain.IGetProductsHandler,
	IGetProductDetailHandler domain.IGetProductDetailHandler,
	IGetProductsByGenderHandler domain.IGetProductsByGenderHandler,
	IGetProductsByCategoryHandler domain.IGetProductsByCategoryHandler,
	IGetProductsPageHandler domain.IGetProductsPageHandler,
) {
	pipeline.RegisterRequestHandler(IGetProductsHandler)
	pipeline.RegisterRequestHandler(IGetProductDetailHandler)
	pipeline.RegisterRequestHandler(IGetProductsByGenderHandler)
	pipeline.RegisterRequestHandler(IGetProductsByCategoryHandler)
	pipeline.RegisterRequestHandler(IGetProductsPageHandler)
}
