package postgres

import (
	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
)

type productRepoImlp struct {
	db *configuration.PostgresCon
}

func NewProductRepoImpl(db *configuration.PostgresCon) domain.IProductRepo {
	return &productRepoImlp{
		db: db,
	}
}
