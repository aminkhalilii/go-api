package services

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
)

type ProductServiceInterface interface {
	getAllProducts() ([]models.Product, error)
	GetUserByID(id int) (*models.User, error)
	CreateProduct(*models.Product) (*models.Product, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

type ProductService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productReposotory repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productReposotory,
	}
}
func (s *ProductService) getAllProducts() ([]models.Product, error) {
	products, err := s.productRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
