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

func NewGetProductDetailHandler(
	repo domain.IProductRepo,
) domain.IGetProductDetailHandler {
	return &handler{
		repo: repo,
	}
}

// Handle implements domain.IGetProductDetailHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetProductDetailRequest) (res *domain.GetProductDetailResponse, err error) {
	logger.Info(ctx, "Get product detail handler start")
	defer logger.Info(ctx, "Get product detail handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	entity, err := h.repo.GetProductDetail(ctx, req.ID)
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	res = &domain.GetProductDetailResponse{
		ID:              entity.ProductID,
		Name:            entity.ProductName,
		Catergory:       entity.ProductCategory,
		Price:           entity.ProductPrice,
		Description:     entity.ProductDescription,
		Image:           entity.ProductImage,
		ProductQuantity: entity.ProductQuantity,
		Gender:          entity.Gender,
	}

	return res, nil

}
