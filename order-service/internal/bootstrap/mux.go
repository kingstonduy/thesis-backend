package configuration

import (
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/order-service/internal/domain"
)

func ResgisterPipeline(
	IExecuteTransactionHandler domain.IExecuteTransactionHandler,
	IGetHistoryHandler domain.IGetHistoryHandler,
) {
	pipeline.RegisterRequestHandler(IExecuteTransactionHandler)
	pipeline.RegisterRequestHandler(IGetHistoryHandler)
}
