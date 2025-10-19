package services

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
)

type UserServiceInterface interface {
	GetAllUsers() []models.User
	GetUserByID(id int) *models.User
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
