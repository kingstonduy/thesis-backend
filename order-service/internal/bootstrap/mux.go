package configuration

import (
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/order-service/internal/domain"
)

func ResgisterPipeline(
	IExecuteTransactionHandler domain.IExecuteTransactionHandler,
	IGetCheckoutItemHandler domain.IGetCheckoutItemHandler,
	IGetHistoryHandler domain.IGetHistoryHandler,
) {
	pipeline.RegisterRequestHandler(IExecuteTransactionHandler)
	pipeline.RegisterRequestHandler(IGetCheckoutItemHandler)
	pipeline.RegisterRequestHandler(IGetHistoryHandler)
}
