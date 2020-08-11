package db

import (
	"context"
	"log"
	"sync"
	"time"
	"github.com/fallenstedt/gin-example/src/config"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


var once sync.Once
var instance *mongo.Client

func Init() *mongo.Client {

	once.Do(func() {
		c := config.Get()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.MongoConnection))

		if err != nil {
			log.Fatalf("Failed to connect to mongo db %f", err)
		}

		instance = client
	})

	return instance
}

func GetCollection(collection string) *mongo.Collection {
	c := config.Get()
	coll := instance.Database(c.DbName).Collection(collection)
	return coll
}