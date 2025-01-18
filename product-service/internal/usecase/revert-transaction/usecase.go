package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
)

type handler struct {
	repo domain.IProductRepo
	db   *configuration.PostgresCon
}

func NewRevertTransactionHandler(
	repo domain.IProductRepo,
	db *configuration.PostgresCon,

) domain.IRevertTransactionHandler {
	return &handler{
		repo: repo,
		db:   db,
	}
}

// Handle implements domain.IRevertTransactionHandler.
func (h *handler) Handle(ctx context.Context, cmd *domain.Command[domain.RevertTransactionRequest]) (res *domain.RevertTransactionResponse, err error) {
	logger.Info(ctx, "Get products handler start")
	defer logger.Info(ctx, "Get products handler end")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	req := cmd.Payload

	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, item := range req.Details {
			product, err := h.repo.GetProductByID(ctx, item.ProductID)
			if err != nil {
				return err
			}

			product.ProductQuantity += item.CartItemQuantity

			conditions := map[string]interface{}{
				"PRODUCT_ID": product.ProductID,
				"UPDATED_AT": product.UpdatedAt,
			}
			columns := map[string]interface{}{
				"PRODUCT_QUANTITY": product.ProductQuantity,
				"UPDATED_AT":       product.UpdatedAt,
			}

			err = h.repo.Update(ctx, columns, conditions)
			if err != nil {
				logger.Error(ctx, err.Error())
				return err
			}
		}

		return nil
	})
	// TODO PUBLISH KAFKA
	if err1 != nil {
		logger.Error(ctx, err1.Error())
		return nil, err1
	}
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return nil, nil
}
