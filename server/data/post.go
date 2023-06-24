package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"suraj.com/refine/models"
)

func getPostColl() *mongo.Collection {
	return getDB().Collection("Posts")
}

func PostGetById(id string) (*models.Post, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	res := getPostColl().FindOne(context.TODO(), filter)

	log.Println("~ PostGetById ~ Result:", res)

	return &models.Post{}, nil
}

func PostDeleteById(id string) (*mongo.DeleteResult, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	res, err := getPostColl().DeleteOne(context.Background(), filter)
	return res, err
}

func PostGetAll() (*[]models.Post, error) {

	options := options.Find()
	res, err := getPostColl().Find(context.TODO(), bson.D{{}}, options)

	log.Println("~ PostGetAll ~ Result:", res)

	return &[]models.Post{}, err
}

func PostCreate(obj *models.Post) (*mongo.InsertOneResult, error) {
	obj.ID = primitive.NewObjectID()
	// obj.UserId = nil
	obj.CreatedAt = time.Now()
	obj.UpdatedAt = time.Now()

	res, err := getPostColl().InsertOne(context.Background(), obj)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func PostDeleteAll() (*mongo.DeleteResult, error) {
	res, err := getPostColl().DeleteMany(context.Background(), bson.D{{}})
	return res, err
}
