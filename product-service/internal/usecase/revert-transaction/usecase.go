package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kingstonduy/go-core/database"
	"github.com/kingstonduy/go-core/errorx"
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
func (h *handler) Handle(ctx context.Context, req *domain.RevertTransactionRequest) (res *domain.RevertTransactionResponse, err error) {
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

			product.ProductQuantity += item.CartItemQuantity
			err = h.repo.UpdateProductByID(ctx, product)

			return err
		}
		return nil
	}, database.WithIsolationLevelOptions(sql.LevelReadCommitted))
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
