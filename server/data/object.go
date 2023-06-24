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

func getObjectColl() *mongo.Collection {
	return getDB().Collection("objects")
}

func ObjectGetById(id string) (*models.Object, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	res := getObjectColl().FindOne(context.TODO(), filter)

	log.Println("~ ObjectGetById ~ Result:", res)

	return &models.Object{}, nil
}

func ObjectDeleteById(id string) (*mongo.DeleteResult, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	res, err := getObjectColl().DeleteOne(context.Background(), filter)
	return res, err
}

func ObjectGetAll() (*[]models.Object, error) {

	options := options.Find()
	res, err := getObjectColl().Find(context.TODO(), bson.D{{}}, options)

	log.Println("~ ObjectGetAll ~ Result:", res)

	return &[]models.Object{}, err
}

func ObjectCreate(obj *models.Object) (*mongo.InsertOneResult, error) {
	obj.ID = primitive.NewObjectID()
	obj.CreatedAt = time.Now()
	obj.UpdatedAt = time.Now()

	res, err := getObjectColl().InsertOne(context.Background(), obj)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ObjectDeleteAll() (*mongo.DeleteResult, error) {
	res, err := getObjectColl().DeleteMany(context.Background(), bson.D{{}})
	return res, err
}
