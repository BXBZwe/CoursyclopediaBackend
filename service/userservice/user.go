package usersvc

import (
	"BackendCoursyclopedia/model"
	userrepo "BackendCoursyclopedia/repository/userrepository"
	"context"
)

type IUserService interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	CreateNewUser(ctx context.Context, user model.User) (*model.User, error)
	DeleteSpecificUser(ctx context.Context, userID string) error
	UpdateSpecificByID(ctx context.Context, userID string, updateUser model.User) (*model.User, error)
}

type UserService struct {
	UserRepository userrepo.IUserRepository
}

func NewUserService(userRepo userrepo.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	return s.UserRepository.FindUserByID(ctx, userID)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.UserRepository.FindAllUsers(ctx)

}

func (s *UserService) CreateNewUser(ctx context.Context, user model.User) (*model.User, error) {
	return s.UserRepository.CreateUser(ctx, user)
}

func (s *UserService) DeleteSpecificUser(ctx context.Context, userID string) error {
	return s.UserRepository.DeleteUserByID(ctx, userID)
}

func (s *UserService) UpdateSpecificByID(ctx context.Context, userID string, updateUser model.User) (*model.User, error) {
	return s.UserRepository.UpdateUserByID(ctx, userID, updateUser)
}
