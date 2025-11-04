package services

import (
	"errors"
	"go-api/internal/models"
	"go-api/internal/repositories"
	"strings"
)

type ProductServiceInterface interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(*models.Product) (*models.Product, error)
	UpdateProduct(id int, product *models.Product) (*models.Product, error)
	DeleteProduct(id int) error
}

type ProductService struct {
	productRepository repositories.ProductRepositoryIntrface
}

func NewProductService(productReposotory repositories.ProductRepositoryIntrface) *ProductService {
	return &ProductService{
		productRepository: productReposotory,
	}
}
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := s.productRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	product, err := s.productRepository.GetProductByID(id)
	if err != nil {
		//should be handled in real world
		return nil, err
	}
	if product == nil {
		return nil, ErrUserNotFound
	}
	return product, nil
}

func (s *ProductService) CreateProduct(user *models.Product) (*models.Product, error) {

	if user.Name == "" || user.Description == "" || user.Price == 0 {
		return nil, errors.New("missing required fields")
	}

	// normalize data
	user.Name = strings.TrimSpace(user.Name)
	user.Description = strings.TrimSpace(strings.ToLower(user.Description))

	return s.productRepository.CreateProduct(user)

}

func (s *ProductService) UpdateProduct(id int, user *models.Product) (*models.Product, error) {
	existing, err := s.productRepository.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrUserNotFound
	}

	user.Description = strings.TrimSpace(strings.ToLower(user.Description))
	user.Name = strings.TrimSpace(user.Name)

	return s.productRepository.UpdateProduct(id, user)
}
func (s *ProductService) DeleteProduct(id int) error {
	existing, err := s.productRepository.GetProductByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrUserNotFound
	}
	return s.productRepository.DeleteProduct(id)
}
