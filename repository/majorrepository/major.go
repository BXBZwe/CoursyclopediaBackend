package majorrepository

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/model/majormodel"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMajorRepository interface {
	CreateMajor(ctx context.Context, majorName string) (string, error)
	DeleteMajor(ctx context.Context, majorId primitive.ObjectID) error
	UpdateMajorName(ctx context.Context, majorId primitive.ObjectID, newName string) error
}

type MajorRepository struct {
	DB *mongo.Client
}

func NewMajorRepository(db *mongo.Client) IMajorRepository {
	return &MajorRepository{
		DB: db,
	}
}

func (r *MajorRepository) CreateMajor(ctx context.Context, majorName string) (string, error) {
	collection := db.GetCollection("majors")
	major := majormodel.Major{
		ID:         primitive.NewObjectID(),
		MajorName:  majorName,
		SubjectIDs: []primitive.ObjectID{},
	}
	_, err := collection.InsertOne(ctx, major)
	if err != nil {
		return "", err
	}
	return major.ID.Hex(), nil
}

func (r *MajorRepository) DeleteMajor(ctx context.Context, majorId primitive.ObjectID) error {
	collection := db.GetCollection("majors")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": majorId})
	return err
}

func (r *MajorRepository) UpdateMajorName(ctx context.Context, majorId primitive.ObjectID, newName string) error {
	collection := db.GetCollection("majors")

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": majorId},
		bson.M{"$set": bson.M{"majorName": newName}},
	)

	return err
}
