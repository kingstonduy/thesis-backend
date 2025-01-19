package configuration

import (
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/order-service/internal/domain"
)

func ResgisterPipeline(
	IExecuteTransactionHandler domain.IExecuteTransactionHandler,
	IGetHistoryHandler domain.IGetHistoryHandler,
	ICartCompletedHandler domain.ICartCompletedHandler,
) {
	pipeline.RegisterRequestHandler(IExecuteTransactionHandler)
	pipeline.RegisterRequestHandler(IGetHistoryHandler)
	pipeline.RegisterRequestHandler(ICartCompletedHandler)
}
