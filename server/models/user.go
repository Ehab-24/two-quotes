package models

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
	LastLogin    time.Time          `json:"lastLogin" bson:"lastLogin"`
	DisplayName  string             `json:"displayName" bson:"displayName"`
	UserName     string             `json:"userName" bson:"userName"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Age          int                `json:"age" bson:"age"`
	NumPosts     int                `json:"numPosts" bson:"numPosts"`
	NumFollowers int                `json:"numFollowers" bson:"numFollowers"`
	NumFollowing int                `json:"numFollowing" bson:"numFollowing"`
}

func NewUser(email string, hashedPass string, userName string, displayName string, age int) *User {
	return &User{
		ID:           primitive.NewObjectID(),
		UserName:     userName,
		DisplayName:  displayName,
		Age:          age,
		Email:        email,
		Password:     hashedPass,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		LastLogin:    time.Now(),
		NumPosts:     0,
		NumFollowers: 0,
		NumFollowing: 0,
	}
}

func (user *User) FromJSON(r *io.ReadCloser) error {
	return json.NewDecoder(*r).Decode(user)
}
