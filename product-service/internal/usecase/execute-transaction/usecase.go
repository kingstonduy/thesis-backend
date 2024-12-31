package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
)

type handler struct {
	repo domain.IProductRepo
	db   *configuration.PostgresCon
}

func NewExecuteTransactionHandler(
	repo domain.IProductRepo,
	db *configuration.PostgresCon,

) domain.IExecuteTransactionHandler {
	return &handler{
		repo: repo,
		db:   db,
	}
}

// Handle implements domain.IExecuteTransactionHandler.
func (h *handler) Handle(ctx context.Context, req *domain.ExecuteTransactionRequest) (res *domain.ExecuteTransactionResponse, err error) {
	logger.Info(ctx, "Get products handler start")
	defer logger.Info(ctx, "Get products handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	var product domain.ProductEntity
	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, item := range req.Details {
			product, err = h.repo.GetProductByID(ctx, item.ProductID)
			if err != nil {
				return err
			}

			if product.ProductQuantity < item.CartItemQuantity {
				err = fmt.Errorf("Product is out of stock")
				return err
			}

			product.ProductQuantity -= item.CartItemQuantity
			err = h.repo.UpdateProductByID(ctx, product)

			return err
		}
		return nil
	})
	if err1 != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	if err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	return res, nil
}
