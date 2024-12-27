package postgres

import (
	"context"

	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
)

type cartRepoImpl struct {
	db *configuration.PostgresCon
}

func NewCartRepo(db *configuration.PostgresCon) domain.ICartRepo {
	return &cartRepoImpl{
		db: db,
	}
}

// AddCartItem implements domain.ICartRepo.
func (c *cartRepoImpl) AddCartItem(ctx context.Context, params domain.AddCartItemParams) error {
	panic("unimplemented")
}

// DeleteCartItem implements domain.ICartRepo.
func (c *cartRepoImpl) DeleteCartItem(ctx context.Context, params domain.DeleteCartItemParams) error {
	panic("unimplemented")
}

// GetCart implements domain.ICartRepo.
func (c *cartRepoImpl) GetCart(ctx context.Context, params domain.GetCartParams) ([]domain.CartItem, error) {
	panic("unimplemented")
}

// UpdateCartItem implements domain.ICartRepo.
func (c *cartRepoImpl) UpdateCartItem(ctx context.Context, params domain.UpdateCartItemParams) error {
	panic("unimplemented")
}
