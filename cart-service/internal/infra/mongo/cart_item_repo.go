package mongo

import (
	"context"

	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type readCartItemRepo struct {
	cfg *configuration.Configuration
	db  *configuration.MongoDB
}

func NewMongoReadCartItemRepo(
	cfg *configuration.Configuration,
	db *configuration.MongoDB,
) domain.IReadCartItemRepo {
	return &readCartItemRepo{
		cfg: cfg,
		db:  db,
	}
}

// Delete implements domain.IReadCartItemRepo.
func (repo *readCartItemRepo) Delete(ctx context.Context, id string) error {
	filter := bson.M{"cart_item_id": id}

	_, err := repo.db.DB.Collection("CART_ITEM").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// Upsert implements domain.IReadCartItemRepo.
func (repo *readCartItemRepo) Upsert(ctx context.Context, entity domain.ReadCartItemEntity) error {
	filter := bson.M{"CART_ITEM_ID": entity.CartItemID}
	update := bson.M{
		"$set": bson.M{
			"USER_ID":            entity.UserID,
			"PRODUCT_ID":         entity.ProductID,
			"CART_ITEM_QUANTITY": entity.CartItemQuantity,
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err := repo.db.DB.Collection("CART_ITEM").UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}
