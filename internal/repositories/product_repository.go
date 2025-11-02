package repositories

import "go-api/internal/models"

type ProductRepository interface {
	getAllProducts() ([]models.Prodcut, error)
	GetUserByID(id int) (*models.User, error)
	CreateProduct(*models.Prodcut) (*models.Prodcut, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}
