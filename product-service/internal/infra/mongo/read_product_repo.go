package mongo

import (
	"context"

	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type mongoReadProductRepo struct {
	db *configuration.MongoCon
}

// NewMongoViewProductRepo initializes a new instance of mongoReadProductRepo.
// Note: It's important to pass in a valid MongoCon instance.
func NewMongoViewProductRepo(con *configuration.MongoCon) domain.IReadProductRepo {
	return &mongoReadProductRepo{
		db: con,
	}
}

// GetAllProduct implements domain.IReadProductRepo.
func (repo *mongoReadProductRepo) GetAllProduct(ctx context.Context) (entities []domain.ReadProductEntity, err error) {
	cursor, err := repo.db.DB.Collection("product").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var entity domain.ReadProductEntity
		if err := cursor.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

// GetProductByCategory implements domain.IReadProductRepo.
func (repo *mongoReadProductRepo) GetProductByCategory(ctx context.Context, category string) ([]domain.ReadProductEntity, error) {
	cursor, err := repo.db.DB.Collection("product").Find(ctx, bson.M{"PRODUCT_CATEGORY": category})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var entities []domain.ReadProductEntity
	for cursor.Next(ctx) {
		var entity domain.ReadProductEntity
		if err := cursor.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

// GetProductByGender implements domain.IReadProductRepo.
func (repo *mongoReadProductRepo) GetProductByGender(ctx context.Context, gender string) ([]domain.ReadProductEntity, error) {
	cursor, err := repo.db.DB.Collection("product").Find(ctx, bson.M{"GENDER": gender})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var entities []domain.ReadProductEntity
	for cursor.Next(ctx) {
		var entity domain.ReadProductEntity
		if err := cursor.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

// GetProductByID implements domain.IReadProductRepo.
func (repo *mongoReadProductRepo) GetProductByID(ctx context.Context, productID string) (domain.ReadProductEntity, error) {
	var entity domain.ReadProductEntity
	err := repo.db.DB.Collection("product").FindOne(ctx, bson.M{"PRODUCT_ID": productID}).Decode(&entity)
	if err != nil {
		return domain.ReadProductEntity{}, err
	}
	return entity, nil
}

// GetProductDetail implements domain.IReadProductRepo.
func (repo *mongoReadProductRepo) GetProductDetail(ctx context.Context, id string) (domain.ReadProductEntity, error) {
	// Here, we assume that the detail view is the same as a lookup by PRODUCT_ID.
	var entity domain.ReadProductEntity
	err := repo.db.DB.Collection("product").FindOne(ctx, bson.M{"PRODUCT_ID": id}).Decode(&entity)
	if err != nil {
		return domain.ReadProductEntity{}, err
	}
	return entity, nil
}
