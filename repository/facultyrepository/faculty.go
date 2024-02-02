package facultyrepository

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/model/facultymodel"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IFacultyRepository interface {
	FindAllFaculties(ctx context.Context) ([]facultymodel.Faculty, error)
	CreateFaculty(ctx context.Context, faculty facultymodel.Faculty) (facultymodel.Faculty, error)
	UpdateFaculty(ctx context.Context, facultyID string, faculty facultymodel.Faculty) (facultymodel.Faculty, error)
	DeleteFaculty(ctx context.Context, facultyID string) error
}

type FacultyRepository struct {
	DB *mongo.Client
}

func NewFacultyRepository(db *mongo.Client) IFacultyRepository {
	return &FacultyRepository{
		DB: db,
	}
}

func (r FacultyRepository) FindAllFaculties(ctx context.Context) ([]facultymodel.Faculty, error) {
	collection := db.GetCollection("faculties")
	var faculties []facultymodel.Faculty

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var faculty facultymodel.Faculty
		if err := cursor.Decode(&faculty); err != nil {
			return nil, err
		}
		faculties = append(faculties, faculty)
	}

	return faculties, nil
}

func (r FacultyRepository) CreateFaculty(ctx context.Context, faculty facultymodel.Faculty) (facultymodel.Faculty, error) {
	collection := db.GetCollection("faculties")
	_, err := collection.InsertOne(ctx, faculty)
	if err != nil {
		return facultymodel.Faculty{}, err
	}
	return faculty, nil
}

func (r FacultyRepository) UpdateFaculty(ctx context.Context, facultyID string, faculty facultymodel.Faculty) (facultymodel.Faculty, error) {
	collection := db.GetCollection("faculties")
	objID, err := primitive.ObjectIDFromHex(facultyID)
	if err != nil {
		return facultymodel.Faculty{}, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": faculty}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return facultymodel.Faculty{}, err
	}
	if result.MatchedCount == 0 {
		return facultymodel.Faculty{}, errors.New("no faculty found with given ID")
	}

	return faculty, nil
}

func (r FacultyRepository) DeleteFaculty(ctx context.Context, facultyID string) error {
	collection := db.GetCollection("faculties")
	objID, err := primitive.ObjectIDFromHex(facultyID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no faculty found with given ID")
	}

	return nil
}
