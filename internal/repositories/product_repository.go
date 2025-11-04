package repositories

import "go-api/internal/models"

type ProductRepositoryIntrface interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(*models.Product) (*models.Product, error)
	UpdateProduct(id int, product *models.Product) (*models.Product, error)
	DeleteProduct(id int) error
}
