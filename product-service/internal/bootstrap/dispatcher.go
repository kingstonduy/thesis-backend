package configuration

import (
	cmd_pipeline "github.com/kingstonduy/go-core/comman-pipeline"
	"github.com/kingstonduy/product-service/internal/domain"
)

func NewDispatcher(
	IExecuteTransactionHandler domain.IExecuteTransactionHandler,
	IRevertTransactionHandler domain.IRevertTransactionHandler,
) cmd_pipeline.DispatcherHandler {
	dp := cmd_pipeline.NewDispatcherCommandHandler()
	dp.RegisterHandler(domain.ORDER_INIT_TRANSACTION_COMMAND, IExecuteTransactionHandler.Handle1)
	dp.RegisterHandler(domain.CART_FAILED_TRANSACTION_COMMAND, IRevertTransactionHandler.Handle1)
	return dp
}
