package postgres

import (
	configuration "github.com/kingstonduy/order-service/internal/bootstrap"
	"github.com/kingstonduy/order-service/internal/domain"
)

type cartRepoImpl struct {
	db *configuration.PostgresCon
}

func NewCartRepo(db *configuration.PostgresCon) domain.IOrderRepo {
	return &cartRepoImpl{
		db: db,
	}
}
