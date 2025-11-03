package repositories

import "go-api/internal/models"

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetUserByID(id int) (*models.User, error)
	CreateProduct(*models.Product) (*models.Product, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}
