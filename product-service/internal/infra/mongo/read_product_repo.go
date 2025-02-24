package mongo

import (
	"context"

	"github.com/kingstonduy/go-core/errorx"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// GetProductByFilter implements domain.IReadProductRepo.
func (repo *mongoReadProductRepo) GetProductByFilter(ctx context.Context, pageNum int, filters map[string]string) (totalPage int, entities []domain.ReadProductEntity, err error) {
	pageSize := repo.cfg.ProductServiceConfig.NumberProductPerPage
	skip := (pageNum - 1) * pageSize

	// Build the query filter
	query := bson.M{}
	if category, ok := filters["CATEGORY"]; ok {
		query["PRODUCT_CATEGORY"] = category
	}
	if gender, ok := filters["GENDER"]; ok {
		query["GENDER"] = gender
	}

	// 1. Count total documents matching the filter
	totalCount, err := repo.db.DB.Collection("product").CountDocuments(ctx, query)
	if err != nil {
		return 0, nil, err
	}

	// Calculate total pages
	totalPage = int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	// 2. Get paginated products
	findOptions := options.Find().
		SetSort(bson.D{{Key: "_id", Value: 1}}).
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize))

	cursor, err := repo.db.DB.Collection("product").Find(ctx, query, findOptions)
	if err != nil {
		return 0, nil, err
	}
	defer cursor.Close(ctx)

	// Decode results
	var products []domain.ReadProductEntity
	if err = cursor.All(ctx, &products); err != nil {
		return 0, nil, err
	}

	if len(products) == 0 {
		return 0, nil, errorx.NotFoundErrorWithDetails("No products found", "")
	}

	return totalPage, products, nil
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
