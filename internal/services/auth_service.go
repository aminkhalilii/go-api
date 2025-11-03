package services

import (
	"errors"
	"go-api/config"
	"go-api/internal/models"
	"go-api/internal/repositories"
	"go-api/pkg/security"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Cache interface {
	GetByID(id int)
}
type AuthServiceInterface interface {
	Register(user *models.User) (*models.User, error)
	Login(email, password string) (*models.User, error)
	GenerateAccessToken(username string) (string, error)
	GenerateRefreshToken(username string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GenerateToken(user *models.User) (map[string]string, error)
}

type AuthService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewAuthService(userRepo repositories.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{userRepo: userRepo}
}
func (s *AuthService) GenerateAccessToken(username string) (string, error) {
	jwtSecret := []byte(config.AppConfig.JwtSecret)

	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (s *AuthService) GenerateRefreshToken(username string) (string, error) {
	jwtSecret := []byte(config.AppConfig.JwtSecret)

	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	jwtSecret := []byte(config.AppConfig.JwtSecret)

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

func (s *AuthService) Register(user *models.User) (*models.User, error) {
	existingUser, _ := s.userRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

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

	err = security.CheckPasswordHash(password, user.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
func (s *AuthService) GenerateToken(user *models.User) (map[string]string, error) {
	accessToken, err := s.GenerateAccessToken(user.Email)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.GenerateRefreshToken(user.Email)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}
