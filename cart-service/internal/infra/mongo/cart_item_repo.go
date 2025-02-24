package mongo

import (
	"context"
	"errors"

	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const cartCollection = "CART_ITEM"

type readCartItemRepo struct {
	cfg *configuration.Configuration
	db  *configuration.MongoDB
}

// NewMongoReadCartItemRepo initializes a new MongoDB repository for cart items
func NewMongoReadCartItemRepo(
	cfg *configuration.Configuration,
	db *configuration.MongoDB,
) domain.IReadCartItemRepo {
	return &readCartItemRepo{
		cfg: cfg,
		db:  db,
	}
}

// Upsert updates an existing cart or inserts a new one
func (r *readCartItemRepo) Upsert(ctx context.Context, entity domain.ReadCartEntity) error {
	collection := r.db.DB.Collection(cartCollection)

	filter := bson.M{"USER_ID": entity.UserID}
	update := bson.M{"$set": entity}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logger.Errorf(ctx, "Error upserting cart in MongoDB: %v", err)
		return err
	}

	logger.Infof(ctx, "Successfully upserted cart for user %s", entity.UserID)
	return nil
}

// Delete removes a specific cart item from the user's cart
func (r *readCartItemRepo) Delete(ctx context.Context, cartItemID string) error {
	collection := r.db.DB.Collection(cartCollection)

	// Pull the item from the cart's items array
	filter := bson.M{"CART_ITEMS.CART_ITEM_ID": cartItemID}
	update := bson.M{"$pull": bson.M{"CART_ITEMS": bson.M{"CART_ITEM_ID": cartItemID}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Errorf(ctx, "Error deleting cart item %s: %v", cartItemID, err)
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("cart item not found")
	}

	logger.Infof(ctx, "Successfully deleted cart item %s", cartItemID)
	return nil
}

// GetCartItemByUserID retrieves a user's cart from MongoDB
func (r *readCartItemRepo) GetCartItemByUserID(ctx context.Context, userID string) (domain.ReadCartEntity, error) {
	collection := r.db.DB.Collection(cartCollection)

	var cart domain.ReadCartEntity
	err := collection.FindOne(ctx, bson.M{"USER_ID": userID}).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		logger.Infof(ctx, "No cart found for user %s", userID)
		return domain.ReadCartEntity{}, err
	} else if err != nil {
		logger.Errorf(ctx, "Error retrieving cart for user %s: %v", userID, err)
		return domain.ReadCartEntity{}, err
	}

	logger.Infof(ctx, "Successfully retrieved cart for user %s", userID)
	return cart, nil
}
