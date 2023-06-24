package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comments are grouped together by post (`ObjectID`) in an array

type Reply struct {
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	Text      string             `json:"text" bson:"text"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	Text      string             `json:"text" bson:"text"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Replies   []Reply            `json:"replies" bson:"replies"`
}

type CommentsDoc struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	ObjectID primitive.ObjectID `json:"ObjectID" bson:"ObjectID"`
	Comments []Comment          `json:"comments" bson:"comments"`
}
