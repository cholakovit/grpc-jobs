package db

import (
	"context"
	"grpc-jobs/server/helper"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client {
	MongoURI := helper.LoadEnv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	mongoDB := helper.LoadEnv("MONGO_DB")
	var collection *mongo.Collection = client.Database(mongoDB).Collection(collectionName)
	return collection
}
