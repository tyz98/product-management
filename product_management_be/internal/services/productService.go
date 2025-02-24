package services

import (
	"product_management/internal/models"
	"product_management/internal/repositories"
)

// ProductService handles the business logic for products
type ProductService struct {
	ProductRepo *repositories.ProductRepository
}

// NewProductService creates a new ProductService instance
func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: repo}
}

// GetPaginatedProducts fetches products with pagination, filtering, and sorting from the database
func (service *ProductService) GetPaginatedProducts(offset, limit int, name, typeStr string, priceMin, priceMax float64, sortBy, sortOrder string) ([]models.Product, int, error) {
	return service.ProductRepo.GetPaginatedProducts(offset, limit, name, typeStr, priceMin, priceMax, sortBy, sortOrder)
}

// CreateProduct creates a new product using the repository
func (service *ProductService) CreateProduct(product *models.Product) error {
	return service.ProductRepo.Create(product)
}

// GetProductByID retrieves a product by ID from the repository
func (service *ProductService) GetProductByID(id uint) (models.Product, error) {
	return service.ProductRepo.GetByID(id)
}

// UpdateProduct updates a product using the repository
func (service *ProductService) UpdateProduct(product *models.Product) error {
	return service.ProductRepo.Update(product)
}

// DeleteProduct deletes a product using the repository
func (service *ProductService) DeleteProduct(id uint) error {
	_, err := service.ProductRepo.GetByID(id)
	if err != nil {
		return err
	}
	return service.ProductRepo.DeleteByID(id)
}
