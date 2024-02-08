package subjectservice

import (
	"BackendCoursyclopedia/model/subjectmodel"
	"BackendCoursyclopedia/repository/majorrepository"
	"BackendCoursyclopedia/repository/subjectrepository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISubjectService interface {
	CreateSubject(ctx context.Context, subject subjectmodel.Subject, majorId string) (string, error)
	DeleteSubject(subjectId string) error
}

type SubjectService struct {
	SubjectRepository subjectrepository.ISubjectRepository
	MajorRepository   majorrepository.IMajorRepository
}

func NewSubjectService(SubjectRepo subjectrepository.ISubjectRepository, MajorRepo majorrepository.IMajorRepository) ISubjectService {
	return &SubjectService{
		SubjectRepository: SubjectRepo,
		MajorRepository:   MajorRepo,
	}
}

func (s *SubjectService) CreateSubject(ctx context.Context, subject subjectmodel.Subject, majorId string) (string, error) {
	subjectId, err := s.SubjectRepository.CreateSubject(ctx, subject)
	if err != nil {
		return "", err
	}

	if subject.SubjectStatus == "" {
		subject.SubjectStatus = "AVAILABLE" // Default status
	}

	// Convert subjectId (primitive.ObjectID) to a hexadecimal string
	subjectIdHex := subjectId.Hex()

	err = s.MajorRepository.AddSubjectToMajor(ctx, majorId, subjectIdHex)
	if err != nil {
		return "", err
	}

	return subjectIdHex, nil
}

func (s *SubjectService) DeleteSubject(subjectId string) error {
	ctx := context.Background()

	// Convert subjectId to ObjectID
	objId, err := primitive.ObjectIDFromHex(subjectId)
	if err != nil {
		return err
	}

	// Delete the subject from the subjects collection
	err = s.SubjectRepository.DeleteSubject(ctx, objId)
	if err != nil {
		return err
	}

	// Remove subject ID from all majors that contain it
	return s.MajorRepository.RemoveSubjectFromMajors(ctx, objId)
}
