package providers

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ProviderHandler() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://hieuls369:hieuls369@cluster0.jdesc9l.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected")
	return client
}

var DB *mongo.Client = ProviderHandler()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("restaurantDB").Collection(collectionName)
}
