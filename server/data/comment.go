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

func CommentCreateOne(postId string, comment *models.Comment) (*mongo.UpdateResult, error) {
	post_oid, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "postId", Value: post_oid}}
	update := bson.D{primitive.E{Key: "push", Value: bson.D{primitive.E{Key: "comments", Value: comment}}}}

	return getCommentsColl().UpdateOne(context.Background(), filter, update)
}

func CommentFindByPostId(pid string) (*models.CommentsDoc, error) {
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

func CommentFindByUserId(uid string) (*[]models.Comment, error) {
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

func CommentDeleteById(postId string, id string) (*mongo.UpdateResult, error) {
	comment_oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post_oid, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "postId", Value: post_oid}}
	update := bson.D{primitive.E{Key: "pull", Value: bson.D{primitive.E{Key: "comments", Value: bson.D{primitive.E{Key: "_id", Value: comment_oid}}}}}}

	return getCommentsColl().UpdateOne(context.Background(), filter, update)
}
