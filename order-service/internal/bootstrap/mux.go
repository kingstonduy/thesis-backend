package configuration

import (
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/order-service/internal/domain"
)

func ResgisterPipeline(
	CheckoutHandler domain.CheckoutHandler,
	GetCheckoutItemHandler domain.GetCheckoutItemHandler,
	GetHistoryHandler domain.GetHistoryHandler,
) {
	pipeline.RegisterRequestHandler(CheckoutHandler)
	pipeline.RegisterRequestHandler(GetCheckoutItemHandler)
	pipeline.RegisterRequestHandler(GetHistoryHandler)
}
