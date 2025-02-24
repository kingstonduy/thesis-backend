package usecase

import (
	"context"
	"fmt"

	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	cartItemRepo domain.IReadCartItemRepo
	productRepo  domain.IReadProductRepo
}

func NewCartItemEventHandler(
	cartItemRepo domain.IReadCartItemRepo,
	productRepo domain.IReadProductRepo,
) domain.ICartItemEventHandler {
	return &handler{
		cartItemRepo: cartItemRepo,
		productRepo:  productRepo,
	}
}

// Handle implements domain.CartItemEventHandler.
func (h *handler) Handle(ctx context.Context, req domain.Event[*domain.CartItemEvent]) (res *domain.CartItemEventRes, err error) {
	logger.Infof(ctx, "CartItemEventHandler start")
	defer logger.Infof(ctx, "CartItemEventHandler end")

	// Recover from panic and log full stack trace
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC CartItemEventHandler: %v", r)
			logger.Errorf(ctx, "PANIC: %v", err)
		}
	}()

	// Handle DELETE operation
	if req.Payload.Deleted {
		err = h.cartItemRepo.Delete(ctx, req.Payload.CartItemID)
		if err != nil {
			logger.Errorf(ctx, "Error deleting cart item: %v", err)
			return nil, err
		}
		return &domain.CartItemEventRes{}, nil
	}

	// Fetch product details
	product, err := h.productRepo.GetProductByID(ctx, req.Payload.ProductID)
	if err != nil {
		logger.Errorf(ctx, "Error fetching product details: %v", err)
		return nil, err
	}

	// Create cart item detail
	cartItemDetail := domain.ReadCartDetail{
		CartItemID:       req.Payload.CartItemID,
		ProductID:        req.Payload.ProductID,
		ProductName:      product.ProductName,
		ProductImage:     product.ProductImage,
		ProductCategory:  product.ProductCategory,
		ProductPrice:     product.ProductPrice,
		Gender:           product.Gender,
		CartItemQuantity: req.Payload.CartItemQuantity,
	}

	// Fetch user cart from MongoDB
	cartItemEntity, err := h.cartItemRepo.GetCartItemByUserID(ctx, req.Payload.UserID)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.Errorf(ctx, "Error retrieving user cart: %v", err)
		return nil, err
	}

	// If cart does not exist, create a new one
	if err == mongo.ErrNoDocuments {
		cartItemEntity = domain.ReadCartEntity{
			UserID: req.Payload.UserID,
			Detail: []domain.ReadCartDetail{cartItemDetail},
		}
		err = h.cartItemRepo.Upsert(ctx, cartItemEntity)
		if err != nil {
			logger.Errorf(ctx, "Error creating new cart: %v", err)
			return nil, err
		}
		return &domain.CartItemEventRes{}, nil
	}

	// Check if item already exists in the cart
	itemUpdated := false
	for i, item := range cartItemEntity.Detail {
		if item.ProductID == req.Payload.ProductID {
			// Update quantity instead of adding a duplicate
			cartItemEntity.Detail[i].CartItemQuantity = req.Payload.CartItemQuantity
			itemUpdated = true
			break
		}
	}

	// If item is new, add it
	if !itemUpdated {
		cartItemEntity.Detail = append(cartItemEntity.Detail, cartItemDetail)
	}

	// Update cart in MongoDB
	err = h.cartItemRepo.Upsert(ctx, cartItemEntity)
	if err != nil {
		logger.Errorf(ctx, "Error updating cart in MongoDB: %v", err)
		return nil, err
	}

	return &domain.CartItemEventRes{}, nil
}
