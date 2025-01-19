package http_server

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/http/fiberx"
	"github.com/kingstonduy/order-service/internal/domain"
	_ "github.com/kingstonduy/order-service/resources/docs"
)

func (s *HttpServer) WithRoutingOption() option {
	return func(s *HttpServer) error {
		s.App.Post("/execute-transaction", s.ExecuteTransaction)
		s.App.Post("/get-history", s.GetHistory)

		return nil
	}
}

//	 	@Tags 			ORDER SERVICE
//		@Summary		ExecuteTransaction
//		@Description	Delete all items on cart -> minus quantity of product in cart service
//		@ID				ExecuteTransaction
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.ExecuteTransactionRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.ExecuteTransactionResponse]			"ok"
//		@Router			/is/v1/order-service/execute-transaction [post]
func (s *HttpServer) ExecuteTransaction(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.ExecuteTransactionRequest, *domain.ExecuteTransactionResponse](ctx)
}

//	 	@Tags 			ORDER SERVICE
//		@Summary		GetPurchasedProducts
//		@Description	Show a history of purchased products from a user
//		@ID				GetPurchasedProducts
//		@Accept			json
//		@Produce		json
//		@Param			request	body		transport.Request[domain.GetHistoryRequest]			false	"Request"
//		@Success		200		{object}	transport.Response[domain.GetHistoryResponse]			"ok"
//		@Router			/is/v1/order-service/get-history [post]
func (s *HttpServer) GetHistory(ctx *fiber.Ctx) error {
	return fiberx.RequestHandlerWithDynamicTimeout[*domain.GetHistoryRequest, *domain.GetHistoryResponse](ctx)
}
