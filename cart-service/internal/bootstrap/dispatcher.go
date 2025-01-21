package configuration

import (
	"github.com/kingstonduy/cart-service/internal/domain"
	cmd_pipeline "github.com/kingstonduy/go-core/comman-pipeline"
)

func NewDispatcher(
	IExecuteTransactionHandler domain.IExecuteTransactionHandler,
) cmd_pipeline.DispatcherHandler {
	dp := cmd_pipeline.NewDispatcherCommandHandler()
	dp.RegisterHandler(domain.PRODUCT_COMPLETED_TRANSACTION_COMMAND, IExecuteTransactionHandler.Handle1)
	return dp
}
