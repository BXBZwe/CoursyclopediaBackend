package facultyservice

import (
	"BackendCoursyclopedia/model/facultymodel"
	facultyrepo "BackendCoursyclopedia/repository/facultyrepository"
	"context"
)

type IFacultyService interface {
	GetAllFaculties(ctx context.Context) ([]facultymodel.Faculty, error)
	CreateFaculty(ctx context.Context, faculty facultymodel.Faculty) (facultymodel.Faculty, error)
	UpdateFaculty(ctx context.Context, facultyID string, faculty facultymodel.Faculty) (facultymodel.Faculty, error)
	DeleteFaculty(ctx context.Context, facultyID string) error
}

type FacultyService struct {
	FacultyRepository facultyrepo.IFacultyRepository
}

func NewFacultyService(facultyRepo facultyrepo.IFacultyRepository) IFacultyService {
	return &FacultyService{
		FacultyRepository: facultyRepo,
	}
}

func (s FacultyService) GetAllFaculties(ctx context.Context) ([]facultymodel.Faculty, error) {
	return s.FacultyRepository.FindAllFaculties(ctx)
}

func (s *FacultyService) CreateFaculty(ctx context.Context, faculty facultymodel.Faculty) (facultymodel.Faculty, error) {
	facultyName := faculty.FacultyName

	return s.FacultyRepository.CreateFaculty(ctx, facultyName)
}

func (s FacultyService) UpdateFaculty(ctx context.Context, facultyID string, faculty facultymodel.Faculty) (facultymodel.Faculty, error) {
	return s.FacultyRepository.UpdateFaculty(ctx, facultyID, faculty)
}

func (s FacultyService) DeleteFaculty(ctx context.Context, facultyID string) error {
	return s.FacultyRepository.DeleteFaculty(ctx, facultyID)
}
