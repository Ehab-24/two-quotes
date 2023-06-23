package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil

func getClient() *mongo.Client {

	if client != nil {
		return client
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var err error

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo DB connected!")
	return client
}

func GetDB() *mongo.Database {
	return getClient().Database("refine")
}
