package models

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ! Since the maximum size limit for a mongodb document is 1MB, a post can only have upto ~32000 likes

type Properties struct {
	TextColor string `json:"textColor" bson:"textColor"`
	BgColor   string `json:"bgColor" bson:"bgColor"`
	BgImage   string `json:"bgImage" bson:"bgImage"`
}

type Like struct {
	UserId primitive.ObjectID `json:"userId" bson:"userId"`
	Type   string             `json:"type" bson:"type"`
}

type Post struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	GroupId      primitive.ObjectID `json:"groupId" bson:"groupId"`
	Type         string             `json:"type" bson:"type"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
	Author       string             `json:"author" bson:"author"`
	Reference    string             `json:"reference" bson:"reference"`
	Properties   Properties         `json:"properties" bson:"properties"`
	Content      string             `json:"content" bson:"content"`
	HideComments bool               `json:"hideComments" bson:"hideComments"`
	Likes        []Like
}

func (obj *Post) FromJSON(r *io.ReadCloser) error {
	return json.NewDecoder(*r).Decode(obj)
}
