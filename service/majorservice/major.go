package majorservice

import (
	"BackendCoursyclopedia/repository/facultyrepository"
	majorrepo "BackendCoursyclopedia/repository/majorrepository"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMajorService interface {
	CreateMajor(majorName string, facultyId string) error
	DeleteMajor(majorId string) error
	UpdateMajor(ctx context.Context, majorId string, newMajorName string, newFacultyId string) error
}

type MajorService struct {
	MajorRepository   majorrepo.IMajorRepository
	FacultyRepository facultyrepository.IFacultyRepository
}

func NewMajorService(MajorRepo majorrepo.IMajorRepository, FacultyRepo facultyrepository.IFacultyRepository) IMajorService {
	return &MajorService{
		MajorRepository:   MajorRepo,
		FacultyRepository: FacultyRepo,
	}
}

func (s *MajorService) CreateMajor(majorName string, facultyId string) error {
	if s.MajorRepository == nil {
		return errors.New("major repository is not initialized")
	}
	if s.FacultyRepository == nil {
		return errors.New("faculty repository is not initialized")
	}
	ctx := context.Background()

	majorId, err := s.MajorRepository.CreateMajor(ctx, majorName)
	if err != nil {
		return err
	}

	return s.FacultyRepository.AddMajorToFaculty(ctx, facultyId, majorId)
}

func (s *MajorService) DeleteMajor(majorId string) error {
	ctx := context.Background()

	objId, err := primitive.ObjectIDFromHex(majorId)
	if err != nil {
		return err
	}

	err = s.MajorRepository.DeleteMajor(ctx, objId)
	if err != nil {
		return err
	}

	return s.FacultyRepository.RemoveMajorFromFaculty(ctx, objId)
}

func (s *MajorService) UpdateMajor(ctx context.Context, majorId string, newMajorName string, newFacultyId string) error {
	majorObjId, err := primitive.ObjectIDFromHex(majorId)
	if err != nil {
		return err
	}

	if newMajorName != "" {
		err = s.MajorRepository.UpdateMajorName(ctx, majorObjId, newMajorName)
		if err != nil {
			return err
		}
	}

	if newFacultyId != "" {
		newFacObjId, err := primitive.ObjectIDFromHex(newFacultyId)
		if err != nil {
			return err
		}

		currentFaculty, err := s.FacultyRepository.FindFacultyByMajorId(ctx, majorObjId)
		if err != nil {
			return err
		}

		if currentFaculty.ID != newFacObjId {
			err = s.FacultyRepository.UpdateFacultyForMajor(ctx, majorObjId, currentFaculty.ID, newFacObjId)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
