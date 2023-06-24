package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"suraj.com/refine/models"
)

func getCommentsColl() *mongo.Collection {
	return getDB().Collection("comments")
}

func CommentsFindByPostId(pid string) (*models.CommentsDoc, error) {
	oid, err := primitive.ObjectIDFromHex(pid)
	if err != nil {
		return nil, err
	}

	var result models.CommentsDoc
	filter := bson.D{primitive.E{Key: "postId", Value: oid}}
	if err = getCommentsColl().FindOne(context.Background(), filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CommentsFindByUserId(uid string) (*[]models.Comment, error) {
	oid, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "postId", Value: oid}}
	cursor, err := getCommentsColl().Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var results []models.Comment
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return &results, nil
}
