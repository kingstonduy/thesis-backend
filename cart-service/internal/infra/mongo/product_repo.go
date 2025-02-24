package mongo

import (
	"context"

	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type mongoReadProductRepo struct {
	cfg *configuration.Configuration
	db  *configuration.MongoDB
}

func NewMongoViewProductRepo(
	con *configuration.MongoDB,
	cfg *configuration.Configuration,
) domain.IReadProductRepo {
	return &mongoReadProductRepo{
		db:  con,
		cfg: cfg,
	}
}

// GetProductByID implements domain.IReadProductRepo.
func (repo *mongoReadProductRepo) GetProductByID(ctx context.Context, productID string) (domain.ReadProductEntity, error) {
	var entity domain.ReadProductEntity
	err := repo.db.DB.Collection("PRODUCT").FindOne(ctx, bson.M{"PRODUCT_ID": productID}).Decode(&entity)
	if err != nil {
		return domain.ReadProductEntity{}, err
	}
	return entity, nil
}
