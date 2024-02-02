package usersvc

import (
	"BackendCoursyclopedia/model"
	userrepo "BackendCoursyclopedia/repository/userrepository"
	"context"
)

type IUserService interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
}

type UserService struct {
	UserRepository userrepo.IUserRepository
}

func NewUserService(userRepo userrepo.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.UserRepository.FindAllUsers(ctx)
}
