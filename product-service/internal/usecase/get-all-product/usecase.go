package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/product-service/internal/domain"
)

type handler struct {
	repo domain.IReadProductRepo
}

func NewGetProductsHandler(
	repo domain.IReadProductRepo,
) domain.IGetProductsHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IGetProductsHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetAllProductRequest) (res *domain.GetAllProductResponse, err error) {
	logger.Info(ctx, "Get products handler start")
	defer logger.Info(ctx, "Get products handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	entities, err := h.repo.GetAllProduct(ctx)
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	products := []domain.GetAllProductResponseDetail{}
	for _, entity := range entities {
		product := domain.GetAllProductResponseDetail{
			ID:              entity.ProductID,
			Name:            entity.ProductName,
			Catergory:       entity.ProductCategory,
			Price:           entity.ProductPrice,
			Description:     entity.ProductDescription,
			Image:           entity.ProductImage,
			ProductQuantity: entity.ProductQuantity,
			Gender:          entity.Gender,
		}
		products = append(products, product)
	}

	res = &domain.GetAllProductResponse{
		Details: products,
	}

	return res, nil
}
