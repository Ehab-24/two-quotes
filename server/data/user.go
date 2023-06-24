package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"suraj.com/refine/models"
)

func getUserColl() *mongo.Collection {
	return getDB().Collection("users")
}

func UserFindById(id string) (*models.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	var user *models.User
	err = getUserColl().FindOne(context.Background(), filter).Decode(user)

	return user, err
}

func UserFindByEmail(email string) (*models.User, error) {

	filter := bson.D{primitive.E{Key: "email", Value: email}}
	var user models.User
	err := getUserColl().FindOne(context.Background(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &user, err
}

func UserCreate(email string, hashedPass string, userName string, displayName string, age int) (*mongo.InsertOneResult, error) {
	user := models.NewUser(email, hashedPass, userName, displayName, age)
	return getUserColl().InsertOne(context.Background(), &user)
}

func UserDeleteById(id string) (*mongo.DeleteResult, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	return getUserColl().DeleteOne(context.Background(), filter)
}

func UserDeleteAll() (*mongo.DeleteResult, error) {
	return getUserColl().DeleteMany(context.Background(), bson.D{{}})
}
