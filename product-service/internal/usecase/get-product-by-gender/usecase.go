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

func NewGetProductsByGenderHandler(
	repo domain.IReadProductRepo,
) domain.IGetProductsByGenderHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IGetProductsByGenderHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetProductsByGenderRequest) (res *domain.GetProductsByGenderResponse, err error) {
	logger.Info(ctx, "GetProductsByGenderHandler start")
	defer logger.Info(ctx, "GetProductsByGenderHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	entities, err := h.repo.GetProductByGender(ctx, req.Gender)
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	products := []domain.GetProductsByGenderResponseDetail{}
	for _, entity := range entities {
		product := domain.GetProductsByGenderResponseDetail{
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

	res = &domain.GetProductsByGenderResponse{
		Details: products,
	}

	return res, nil
}
