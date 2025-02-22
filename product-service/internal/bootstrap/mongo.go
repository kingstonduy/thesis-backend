package configuration

import (
	"context"
	"fmt"
	"log"

	"github.com/kingstonduy/go-core/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCon struct {
	DB *mongo.Database
}

func NewMongoCon(cfg *Configuration) *MongoCon {
	c := cfg.MongoConfig

	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d/",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
	)

	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Ping the database to check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	logger.Info(context.Background(), "Connected to mongoDB")
	db := client.Database(c.Database)
	return &MongoCon{
		DB: db,
	}
}
