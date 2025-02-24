package repositories

import (
	"github.com/jinzhu/gorm"
	"product_management/internal/models"
)

// ProductRepository handles database operations for products
type ProductRepository struct {
	DB *gorm.DB
}

// NewProductRepository creates a new ProductRepository instance
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// GetPaginatedProducts retrieves products from the database with pagination, filtering, and sorting
func (repo *ProductRepository) GetPaginatedProducts(offset, limit int, name, typeStr string, priceMin, priceMax float64, sortBy, sortOrder string) ([]models.Product, int, error) {
	var products []models.Product
	var total int
	query := repo.DB.Model(&models.Product{})

	// Apply fuzzy search for name and type if provided
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if typeStr != "" {
		query = query.Where("type LIKE ?", "%"+typeStr+"%")
	}

	// Apply price range search if provided
	if priceMin > 0 {
		query = query.Where("price >= ?", priceMin)
	}
	if priceMax > 0 {
		query = query.Where("price <= ?", priceMax)
	}

	// Get the total count of products matching the filters (no pagination)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting if specified
	if sortBy != "" {
		if sortOrder == "desc" {
			query = query.Order(sortBy + " desc")
		} else {
			query = query.Order(sortBy + " asc")
		}
	}

	// Execute the query and retrieve the results
	if err := query.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, total, err
	}

	return products, total, nil
}

// Create inserts a new product into the database
func (repo *ProductRepository) Create(product *models.Product) error {
	return repo.DB.Create(product).Error
}

// GetByID retrieves a product by ID from the database
func (repo *ProductRepository) GetByID(id uint) (models.Product, error) {
	var product models.Product
	if err := repo.DB.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

// DeleteByID deletes a product by ID from the database
func (repo *ProductRepository) DeleteByID(id uint) error {
	var product models.Product
	return repo.DB.Delete(&product, id).Error
}

// Update updates an existing product in the database
func (repo *ProductRepository) Update(product *models.Product) error {
	return repo.DB.Save(product).Error
}
