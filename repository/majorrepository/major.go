package majorrepository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IMajorRepository interface {
}

type MajorRepository struct {
	DB *mongo.Client
}

func NewMajorRepository(db *mongo.Client) IMajorRepository {
	return &MajorRepository{
		DB: db,
	}
}
