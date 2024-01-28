package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Email       string               `bson:"email"`
	Role        string               `bson:"role"`
	Wishlists   []primitive.ObjectID `bson:"wishlists"`
	PhoneNumber string               `bson:"phoneNumber"`
	Profile     struct {
		FirstName string `bson:"firstName"`
		LastName  string `bson:"lastName"`
	} `bson:"profile"`
	Faculty primitive.ObjectID `bson:"faculty,omitempty"`
}
