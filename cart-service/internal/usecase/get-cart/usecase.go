package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
)

type handler struct {
	repo domain.IReadCartItemRepo
}

func NewGetCartHandler(
	repo domain.IReadCartItemRepo,
) domain.IGetCartHandler {
	return &handler{
		repo: repo,
	}
}

// Handle retrieves cart details for a user.
func (h *handler) Handle(ctx context.Context, req *domain.GetCartRequest) (res *domain.GetCartResponse, err error) {
	logger.Info(ctx, "IGetCartHandler start")
	defer logger.Info(ctx, "IGetCartHandler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC IGetCartHandler: %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	// Fetch cart from MongoDB
	cart, err := h.repo.GetCartItemByUserID(ctx, req.UserID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			// Return empty cart instead of an error
			return &domain.GetCartResponse{CartItems: []domain.GetCartItemDetail{}}, nil
		}
		errx := errorx.OutboundErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	// Convert MongoDB schema (ReadCartDetail) to response struct (GetCartItemDetail)
	var getCartItemDetails []domain.GetCartItemDetail
	for _, item := range cart.Detail {
		getCartItemDetails = append(getCartItemDetails, domain.GetCartItemDetail{
			CartItemID:       item.CartItemID,
			ProductID:        item.ProductID,
			ProductName:      item.ProductName,
			ProductImage:     item.ProductImage,
			ProductCatergory: item.ProductCategory,
			ProductPrice:     item.ProductPrice,
			CartItemQuantity: item.CartItemQuantity,
		})
	}

	// Construct response
	res = &domain.GetCartResponse{
		CartItems: getCartItemDetails,
	}

	return res, nil
}
