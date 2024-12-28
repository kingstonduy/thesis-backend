package domain

import (
	"github.com/shopspring/decimal"
)

type GetAllProductRequest struct {
}

type GetAllProductResponse struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	ImageURL      string          `json:"imageUrl"`
	Price         decimal.Decimal `json:"price"`
	AverageRating decimal.Decimal `json:"averageRating"`
}
