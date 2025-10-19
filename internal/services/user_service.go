package services

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
)

type UserServiceInterface interface {
	GetAllUsers() []models.User
	GetUserByID(id int) *models.User
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

func (s *UserService) GetAllUsers() []models.User {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		// در حالت واقعی بهتره مدیریت خطا انجام بدیم
		return []models.User{}
	}
	return users
}
func (s *UserService) GetUserByID(id int) *models.User {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		// در حالت واقعی بهتره مدیریت خطا انجام بدیم
		return nil
	}
	return user
}
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.userRepository.CreateUser(user)

}

func (s *UserService) UpdateUser(id int, user *models.User) (*models.User, error) {
	return s.userRepository.UpdateUser(id, user)
}
func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}
