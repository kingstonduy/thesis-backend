package mongo

import (
	"context"

	"github.com/kingstonduy/go-core/errorx"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoReadProductRepo struct {
	cfg *configuration.Configuration
	db  *configuration.MongoCon
}

// NewMongoViewProductRepo initializes a new instance of mongoReadProductRepo.
// Note: It's important to pass in a valid MongoCon instance.
func NewMongoViewProductRepo(
	con *configuration.MongoCon,
	cfg *configuration.Configuration,
) domain.IReadProductRepo {
	return &mongoReadProductRepo{
		db:  con,
		cfg: cfg,
	}
}

// GetAllProductPage implements domain.IReadProductRepo.
func (repo *mongoReadProductRepo) GetAllProductPage(ctx context.Context, pageNum int) (totalPage int, entities []domain.ReadProductEntity, err error) {
	pageSize := repo.cfg.ProductServiceConfig.NumberProductPerPage
	skip := (pageNum - 1) * pageSize

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$facet", Value: bson.D{
			{Key: "products", Value: bson.A{
				bson.D{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}}, // Sort by _id
				bson.D{{Key: "$skip", Value: skip}},                           // Skip documents
				bson.D{{Key: "$limit", Value: pageSize}},                      // Limit documents
			}},
			{Key: "totalCount", Value: bson.A{
				bson.D{{Key: "$count", Value: "count"}}, // Count total products
			}},
		}}},
	}

	cursor, err := repo.db.DB.Collection("product").Aggregate(ctx, pipeline)
	if err != nil {
		return 0, nil, err
	}
	defer cursor.Close(ctx)

	// Define a struct to decode the aggregated response properly
	var result struct {
		Products   []domain.ReadProductEntity `bson:"products"`
		TotalCount []struct {
			Count int `bson:"count"`
		} `bson:"totalCount"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, nil, err
		}
	}

	// Extract total count safely
	if len(result.TotalCount) > 0 {
		totalPage = (result.TotalCount[0].Count + pageSize - 1) / pageSize
	}

	if len(result.Products) == 0 {
		return 0, nil, errorx.NotFoundErrorWithDetails("No products found", "")
	}

	return totalPage, result.Products, nil
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
