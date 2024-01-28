package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Major struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	MajorName string             `bson:"majorName"`
	Subjects  []Subject          `bson:"subjects"`
}

type Faculty struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FacultyName string             `bson:"facultyName"`
	Majors      []Major            `bson:"majors"`
}
