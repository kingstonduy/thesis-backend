package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/thesis-backend/internal/domain"
	"github.com/kingstonduy/thesis-backend/internal/pkg/utils"
)

type handler struct {
	repo            domain.IProductRepo
	commentOutbound domain.ICommentOutbound
}

func NewGetProductsHandler(
	repo domain.IProductRepo,
	commentOutbound domain.ICommentOutbound,
) domain.IGetProductsHandler {
	return &handler{
		repo:            repo,
		commentOutbound: commentOutbound,
	}
}

// Handle implements domain.IGetProductsHandler.
func (h *handler) Handle(ctx context.Context, req *domain.GetProductsRequest) (res *domain.GetProductsResponse, err error) {
	logger.Info(ctx, "Get products handler start")
	defer logger.Info(ctx, "Get products handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	trace := transport.GetTraceByCtx(ctx)

	entities, err := h.repo.GetAllProduct(ctx)
	if err != nil {
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	// TODO call rest to get avg rating

	products := []domain.Product{}
	for _, entity := range entities {

		commentRes, err := h.commentOutbound.GetAvgerageRatingByProductID(ctx, domain.GetAvgerageRatingByProductIDRequest{ProductID: entity.ProductID}, utils.GenRequestTrace(trace, "comment-service", ""))
		if err != nil {
			errx := errorx.OutboundErrorWithDetails(err.Error(), "")
			logger.Error(ctx, errx.Error())
			return nil, errx
		}

		product := domain.Product{
			ID:            entity.ProductID,
			Name:          entity.ProductName,
			ImageURL:      entity.ProductImage,
			Price:         entity.ProductPrice,
			AverageRating: commentRes.Rating,
		}
		products = append(products, product)
	}

	res = &domain.GetProductsResponse{
		Products: products,
	}

	return res, nil
}
