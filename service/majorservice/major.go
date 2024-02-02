package majorservice

import (
	"BackendCoursyclopedia/repository/facultyrepository"
	majorrepo "BackendCoursyclopedia/repository/majorrepository"
)

type IMajorService interface {
}

type MajorService struct {
	MajorRepository   majorrepo.IMajorRepository
	FacultyRepository facultyrepository.IFacultyRepository
}

func NewMajorService(MajorRepo majorrepo.IMajorRepository) IMajorService {
	return &MajorService{MajorRepository: MajorRepo}
}
