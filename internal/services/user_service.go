package services

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
	"go-api/pkg/security"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(*models.User) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepository repositories.UserRepositoryInterface) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		//should be handled in real world
		return nil, err
	}
	return users, nil
}
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		//should be handled in real world
		return nil, err
	}
	return user, nil
}
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	hashed, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed
	return s.userRepository.CreateUser(user)

}

func (s *UserService) UpdateUser(id int, user *models.User) (*models.User, error) {
	return s.userRepository.UpdateUser(id, user)
}
func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}
