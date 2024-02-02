package usersvc

import (
	"BackendCoursyclopedia/model/usermodel"
	userrepo "BackendCoursyclopedia/repository/userrepository"
	"context"
)

type IUserService interface {
	GetAllUsers(ctx context.Context) ([]usermodel.User, error)
}

type UserService struct {
	UserRepository userrepo.IUserRepository
}

func NewUserService(userRepo userrepo.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]usermodel.User, error) {
	return s.UserRepository.FindAllUsers(ctx)
}
