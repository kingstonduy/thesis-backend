package configuration

import (
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/product-service/internal/domain"
)

func ResgisterPipeline(
	IGetProductsHandler domain.IGetProductsHandler,
	IGetProductDetailHandler domain.IGetProductDetailHandler,
	IExecuteTransactionHandler domain.IExecuteTransactionHandler,
	IRevertTransactionHandler domain.IRevertTransactionHandler,
	IGetProductsByGenderHandler domain.IGetProductsByGenderHandler,
	IGetProductsByCategoryHandler domain.IGetProductsByCategoryHandler,

) {
	pipeline.RegisterRequestHandler(IGetProductsHandler)
	pipeline.RegisterRequestHandler(IGetProductDetailHandler)
	pipeline.RegisterRequestHandler(IExecuteTransactionHandler)
	pipeline.RegisterRequestHandler(IRevertTransactionHandler)
	pipeline.RegisterRequestHandler(IGetProductsByGenderHandler)
	pipeline.RegisterRequestHandler(IGetProductsByCategoryHandler)
}
