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
	FindAllSubjects(ctx context.Context) ([]subjectmodel.Subject, error)
	FindSubjectbyID(ctx context.Context, subjectId string) (*subjectmodel.Subject, error)
	FindSubjectsByIDs(ctx context.Context, subjectIDs []primitive.ObjectID) ([]subjectmodel.Subject, error)
	CreateSubject(ctx context.Context, subject subjectmodel.Subject) (primitive.ObjectID, error)
	DeleteSubject(ctx context.Context, subjectId primitive.ObjectID) error
	UpdateSubject(ctx context.Context, subjectId primitive.ObjectID, updates bson.M) error
}

type SubjectRepository struct {
	DB *mongo.Client
}

func NewSubjectRepository(db *mongo.Client) ISubjectRepository {
	return &SubjectRepository{DB: db}
}

func (r SubjectRepository) FindAllSubjects(ctx context.Context) ([]subjectmodel.Subject, error) {
	collection := db.GetCollection("subjects")
	var subjects []subjectmodel.Subject

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var subject subjectmodel.Subject
		if err := cursor.Decode(&subject); err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (r *SubjectRepository) FindSubjectbyID(ctx context.Context, subjectId string) (*subjectmodel.Subject, error) {
	collection := db.GetCollection("subjects")
	var subject subjectmodel.Subject

	objID, err := primitive.ObjectIDFromHex(subjectId)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&subject)
	if err != nil {
		return nil, err
	}
	return &subject, nil
}

func (r *SubjectRepository) FindSubjectsByIDs(ctx context.Context, subjectIDs []primitive.ObjectID) ([]subjectmodel.Subject, error) {
	collection := db.GetCollection("subjects")
	var subjects []subjectmodel.Subject

	filter := bson.M{"_id": bson.M{"$in": subjectIDs}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var subject subjectmodel.Subject
		if err := cursor.Decode(&subject); err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return subjects, nil
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

func (r *SubjectRepository) UpdateSubject(ctx context.Context, subjectId primitive.ObjectID, updates bson.M) error {
	collection := db.GetCollection("subjects")

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": subjectId},
		bson.M{"$set": updates},
	)

	return err
}
