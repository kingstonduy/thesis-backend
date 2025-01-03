package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/product-service/internal/domain"
)

type handler struct {
	repo domain.IProductRepo
}

func NewGetProductsByCategoryHandler(
	repo domain.IProductRepo,
) domain.IGetProductsByCategoryHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IGetProductsByCategoryHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetProductsByCategoryRequest) (res *domain.GetProductsByCategoryResponse, err error) {
	logger.Info(ctx, "GetProductsByCategoryHandler start")
	defer logger.Info(ctx, "GetProductsByCategoryHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	entities, err := h.repo.GetProductByCategory(ctx, req.Category)
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	products := []domain.GetProductsByCategoryResponseDetail{}
	for _, entity := range entities {
		product := domain.GetProductsByCategoryResponseDetail{
			ID:              entity.ProductID,
			Name:            entity.ProductName,
			ImageURL:        entity.ProductImage,
			Price:           entity.ProductPrice,
			AverageRating:   entity.AvgRating,
			ProductQuantity: entity.ProductQuantity,
			TotalRating:     entity.TotalRating,
		}
		products = append(products, product)
	}

	res = &domain.GetProductsByCategoryResponse{
		Details: products,
	}

	return res, nil
}
