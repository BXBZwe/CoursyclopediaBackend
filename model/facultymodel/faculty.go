package facultymodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Faculty struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	FacultyName string               `bson:"facultyName"`
	MajorIDs    []primitive.ObjectID `bson:"majorIDs"`
}
