package repositories

import "go-api/internal/models"

// UserRepository defines data access behavior
type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(*models.User) (*models.User, error)
}
