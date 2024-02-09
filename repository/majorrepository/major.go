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
	AddSubjectToMajor(ctx context.Context, majorId string, subjectId string) error
	RemoveSubjectFromMajors(ctx context.Context, subjectId primitive.ObjectID) error
	FindMajorBySubjectId(ctx context.Context, subjectId primitive.ObjectID) (majormodel.Major, error)
	UpdatemajorforSubject(ctx context.Context, subjectId primitive.ObjectID, currentmajorId primitive.ObjectID, newmajorId primitive.ObjectID) error
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

func (r *MajorRepository) AddSubjectToMajor(ctx context.Context, majorId string, subjectId string) error {
	collection := db.GetCollection("majors")

	mid, err := primitive.ObjectIDFromHex(majorId)
	if err != nil {
		return err
	}
	sid, err := primitive.ObjectIDFromHex(subjectId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": mid}
	update := bson.M{"$addToSet": bson.M{"subjectIDs": sid}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *MajorRepository) RemoveSubjectFromMajors(ctx context.Context, subjectId primitive.ObjectID) error {
	collection := db.GetCollection("majors")

	// Update all majors that contain the subjectId to remove it
	_, err := collection.UpdateMany(
		ctx,
		bson.M{"subjectIDs": subjectId}, // Query for majors that contain the subjectId
		bson.M{"$pull": bson.M{"subjectIDs": subjectId}}, // Remove the subjectId from the subjects list
	)
	return err
}

func (r *MajorRepository) FindMajorBySubjectId(ctx context.Context, subjectId primitive.ObjectID) (majormodel.Major, error) {
	collection := db.GetCollection("majors")
	var major majormodel.Major

	filter := bson.M{"subjectIDs": subjectId}
	err := collection.FindOne(ctx, filter).Decode(&major)
	if err != nil {
		return majormodel.Major{}, err
	}

	return major, nil
}

func (r *MajorRepository) UpdatemajorforSubject(ctx context.Context, subjectId primitive.ObjectID, currentmajorId primitive.ObjectID, newmajorId primitive.ObjectID) error {
	collection := db.GetCollection("majors")

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": currentmajorId},
		bson.M{"$pull": bson.M{"subjectIDs": subjectId}},
	)

	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": newmajorId},
		bson.M{"$addtoSet": bson.M{"subjectIDs": subjectId}},
	)

	return err
}
