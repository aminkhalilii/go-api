package services

import (
	"errors"
	"go-api/internal/models"
	"go-api/internal/repositories"
	"go-api/pkg/security"
	"strings"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
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
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		//should be handled in real world
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {

	if user.Email == "" || user.Password == "" || user.Name == "" {
		return nil, errors.New("missing required fields")
	}

	// normalize data
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Name = strings.TrimSpace(user.Name)

	//check if user exists
	existUser, err := s.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existUser != nil {
		return nil, ErrEmailExists
	}
	//hash password
	hashed, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed
	return s.userRepository.CreateUser(user)

}

func (s *UserService) UpdateUser(id int, user *models.User) (*models.User, error) {
	existing, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrUserNotFound
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Name = strings.TrimSpace(user.Name)

	existingByEmail, err := s.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existingByEmail != nil && existingByEmail.ID != id {
		return nil, ErrEmailExists
	}

	return s.userRepository.UpdateUser(id, user)
}
func (s *UserService) DeleteUser(id int) error {
	existing, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrUserNotFound
	}
	return s.userRepository.DeleteUser(id)
}
