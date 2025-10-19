package repositories

import "go-api/internal/models"

// UserRepository defines data access behavior
type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, error)
}

// UserRepositoryImpl is a concrete implementation
type UserRepositoryImpl struct{ name string }

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{name: "amin"}
}

func (r *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	users := []models.User{
		{ID: 1, Name: "amin", Email: "aminkhalili@gmail.com", Password: "12345"},
		{ID: 2, Name: "sara", Email: "sara@example.com", Password: "54321"},
	}
	return users, nil
}
