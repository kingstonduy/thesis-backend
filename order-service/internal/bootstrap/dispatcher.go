package configuration

import (
	cmd_pipeline "github.com/kingstonduy/go-core/comman-pipeline"
	"github.com/kingstonduy/order-service/internal/domain"
)

func NewDispatcher(
	ICartCompletedHandler domain.ICartCompletedHandler,
	IRevertTransactionHandler domain.IRevertTransactionHandler,
) cmd_pipeline.DispatcherHandler {
	dp := cmd_pipeline.NewDispatcherCommandHandler()
	dp.RegisterHandler(domain.CART_COMPLETED_TRANSACTION_COMMAND, ICartCompletedHandler.Handle1)
	dp.RegisterHandler(domain.CART_FAILED_TRANSACTION_COMMAND, IRevertTransactionHandler.Handle1)
	dp.RegisterHandler(domain.PRODUCT_FAILED_TRANSACTION_COMMAND, IRevertTransactionHandler.Handle1)
	return dp
}
