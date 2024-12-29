package usecase

import "github.com/kingstonduy/order-service/internal/domain"

func getCartExecuteTransactionRequest(req domain.ExecuteTransactionRequest) domain.CartExecuteTransactionRequest {
	var res = domain.CartExecuteTransactionRequest{}

	for _, reqDetails := range req.Details {
		res.Details = append(res.Details, domain.CartExecuteTransactionRequestDetails{
			CartItemID: reqDetails.CartItemID,
		})
	}

	return res
}

func getCartRevertTransactionoutRequest(req domain.ExecuteTransactionRequest) domain.CartRevertTransactionRequest {
	var res = domain.CartRevertTransactionRequest{}

	res.UserID = req.UserID

	for _, reqDetails := range req.Details {
		res.Details = append(res.Details, domain.CartRevertTransactionRequestDetails{
			ProductID:        reqDetails.ProductID,
			CartItemQuantity: reqDetails.CartItemQuantity,
		})
	}

	return res
}

func getProductExecuteTransactionRequest(req domain.ExecuteTransactionRequest) domain.ProductExecuteTransactionRequest {
	var res = domain.ProductExecuteTransactionRequest{}

	for _, reqDetails := range req.Details {
		res.Details = append(res.Details, domain.ProductExecuteTransactionRequestDetails{
			ProductID:       reqDetails.ProductID,
			ProductQuantity: reqDetails.CartItemQuantity,
		})
	}

	return res
}

func getProductRevertTransactionRequest(req domain.ExecuteTransactionRequest) domain.ProductRevertTransactionRequest {
	var res = domain.ProductRevertTransactionRequest{}

	for _, reqDetails := range req.Details {
		res.Details = append(res.Details, domain.ProductRevertTransactionRequestDetails{
			ProductID:       reqDetails.ProductID,
			ProductQuantity: reqDetails.CartItemQuantity,
		})
	}

	return res
}
