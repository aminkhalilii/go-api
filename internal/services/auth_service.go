package services

import (
	"errors"
	"go-api/config"
	"go-api/internal/models"
	"go-api/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Cache interface{
	GetByID(id int)
	
}
type AuthServiceInterface interface {
	Register(user *models.User) (*models.User, error)
	Login(email, password string) (*models.User, error)
	GenerateToken(user *models.User) (string, error)
}

type AuthService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewAuthService(userRepo repositories.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(user *models.User) (*models.User, error) {
	existingUser, _ := s.userRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	newUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *AuthService) GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(config.AppConfig.JwtSecret))
}
