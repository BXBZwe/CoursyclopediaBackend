package usermodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Email       string               `bson:"email"`
	Roles       []string             `bson:"roles"`
	Wishlists   []primitive.ObjectID `bson:"wishlists"`
	PhoneNumber string               `bson:"phoneNumber"`
	Profile     struct {
		FirstName string `bson:"firstName"`
		LastName  string `bson:"lastName"`
	} `bson:"profile"`
	FacultyID primitive.ObjectID `bson:"facultyId,omitempty"`
}
