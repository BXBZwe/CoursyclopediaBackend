package subjectrepository

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/model/subjectmodel"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ISubjectRepository interface {
	CreateSubject(ctx context.Context, subject subjectmodel.Subject) (primitive.ObjectID, error)
	DeleteSubject(ctx context.Context, subjectId primitive.ObjectID) error
}

type SubjectRepository struct {
	DB *mongo.Client
}

func NewSubjectRepository(db *mongo.Client) ISubjectRepository {
	return &SubjectRepository{DB: db}
}

func (r *SubjectRepository) CreateSubject(ctx context.Context, subject subjectmodel.Subject) (primitive.ObjectID, error) {
	collection := db.GetCollection("subjects")
	result, err := collection.InsertOne(ctx, subject)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *SubjectRepository) DeleteSubject(ctx context.Context, subjectId primitive.ObjectID) error {
	collection := db.GetCollection("subjects")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": subjectId})
	return err
}

// Updates the details of a specific subject.
