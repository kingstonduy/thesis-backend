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

func NewGetProductsPageHandler(
	repo domain.IReadProductRepo,
) domain.IGetProductsPageHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IGetProductsPageHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetAllProductPageRequest) (res *domain.GetAllProductPageResponse, err error) {
	logger.Info(ctx, "Get products handler start")
	defer logger.Info(ctx, "Get products handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	filter := map[string]string{}
	if req.Category != "" {
		filter["CATEGORY"] = req.Category
	}
	if req.Gender != "" {
		filter["GENDER"] = req.Gender
	}

	totalPage, entities, err := h.repo.GetProductByFilter(ctx, req.PageNumber, filter)
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	products := []domain.GetAllProductPageResponseDetail{}
	for _, entity := range entities {
		product := domain.GetAllProductPageResponseDetail{
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

	res = &domain.GetAllProductPageResponse{
		Details:   products,
		TotalPage: totalPage,
	}

	return res, nil
}
