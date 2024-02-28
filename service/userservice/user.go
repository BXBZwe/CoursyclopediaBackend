package usersvc

import (
	"BackendCoursyclopedia/model/usermodel"
	userrepo "BackendCoursyclopedia/repository/userrepository"
	"context"
)

type IUserService interface {
	GetAllUsers(ctx context.Context) ([]usermodel.User, error)
	GetUserByID(ctx context.Context, userID string) (*usermodel.User, error)
	GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error)
	CreateNewUser(ctx context.Context, user usermodel.User) (*usermodel.User, error)
	DeleteSpecificUser(ctx context.Context, userID string) error
	UpdateSpecificByID(ctx context.Context, userID string, updateUser usermodel.User) (*usermodel.User, error)
}

type UserService struct {
	UserRepository userrepo.IUserRepository
}

func NewUserService(userRepo userrepo.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*usermodel.User, error) {
	return s.UserRepository.FindUserByID(ctx, userID)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]usermodel.User, error) {
	return s.UserRepository.FindAllUsers(ctx)

}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	return s.UserRepository.GetUserByEmail(ctx, email)

}

func (s *UserService) CreateNewUser(ctx context.Context, user usermodel.User) (*usermodel.User, error) {
	return s.UserRepository.CreateUser(ctx, user)
}

func (s *UserService) DeleteSpecificUser(ctx context.Context, userID string) error {
	return s.UserRepository.DeleteUserByID(ctx, userID)
}

func (s *UserService) UpdateSpecificByID(ctx context.Context, userID string, updateUser usermodel.User) (*usermodel.User, error) {
	return s.UserRepository.UpdateUserByID(ctx, userID, updateUser)
}
