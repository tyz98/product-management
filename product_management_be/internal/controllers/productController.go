package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"product_management/internal/models"
	"product_management/internal/services"
	"product_management/pkg/utils"
	"strconv"
	"time"
)

// ProductController handles HTTP requests for product operations
type ProductController struct {
	ProductService *services.ProductService
}

// NewProductController creates a new ProductController instance
func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}

// GetProducts handles the request to get all products with pagination, filtering, and sorting
func (controller *ProductController) GetProducts(c *gin.Context) {
	// Get page and pageSize query parameters from the URL
	pageStr := c.DefaultQuery("page", "1")           // Default to page 1 if not provided
	pageSizeStr := c.DefaultQuery("page_size", "10") // Default to page_size 10 if not provided

	// Get name, type, and price query parameters for filtering
	name := c.DefaultQuery("name", "")             // Default to empty string if not provided
	typeStr := c.DefaultQuery("type", "")          // Default to empty string if not provided
	priceMinStr := c.DefaultQuery("price_min", "") // Default to empty string if not provided
	priceMaxStr := c.DefaultQuery("price_max", "") // Default to empty string if not provided

	// Get sort query parameters for sorting
	sortBy := c.DefaultQuery("sort_by", "id")        // Default to sorting by id if not provided
	sortOrder := c.DefaultQuery("sort_order", "asc") // Default to ascending order if not provided
	if sortBy == "" {
		sortBy = "id"
	}
	if sortOrder == "" {
		sortOrder = "asc"
	}

	// Convert page and pageSize to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize number"})
		return
	}

	// Convert price range values to float64 (if present)
	var priceMin, priceMax float64
	if priceMinStr != "" {
		priceMin, err = strconv.ParseFloat(priceMinStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price_min value"})
			return
		}
	}
	if priceMaxStr != "" {
		priceMax, err = strconv.ParseFloat(priceMaxStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price_max value"})
			return
		}
	}

	// Calculate the offset based on the page and pageSize
	offset := (page - 1) * pageSize

	// Fetch the filtered, sorted, and paginated products from the service layer
	products, total, err := controller.ProductService.GetPaginatedProducts(offset, pageSize, name, typeStr, priceMin, priceMax, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the paginated products and total count
	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"total":    total,
	})
}

// CreateProduct handles the request to create a new product
func (controller *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind the incoming JSON data to the product struct
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate product fields before processing the image
	if product.Name == "" || product.Type == "" || product.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, Type, and Price must be valid"})
		return
	}

	// Create the product in the database
	if err := controller.ProductService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProduct handles the request to get a product by ID
func (controller *ProductController) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	product, err := controller.ProductService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles the request to update a product by ID
func (controller *ProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	var product models.Product
	if err = c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(id)

	if err = controller.ProductService.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct handles the request to delete a product by ID
func (controller *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	if err = controller.ProductService.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete product"})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UploadImage handles the image upload and returns the image URL
func (controller *ProductController) UploadImage(c *gin.Context) {
	// Check if an image is provided in the form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	// Validate image file type
	if !utils.IsValidImage(file) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}

	// Save the image and get the file path
	// Define a directory to save uploaded files
	uploadDir := "/static"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		os.Mkdir(uploadDir, os.ModePerm)
	}
	filePath := filepath.Join(uploadDir, time.Now().Format("20060102_150405_")+file.Filename)
	err = utils.SaveImage(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
		return
	}

	// Return the image URL
	c.JSON(http.StatusOK, gin.H{"url": filePath})
}

// GetProductTypes handles the request to get all allowed product types
func (controller *ProductController) GetProductTypes(c *gin.Context) {
	types := []string{"Electronics", "Clothing", "Beauty", "Books", "Jewelry", "Furniture", "Grocery"}
	c.JSON(http.StatusOK, gin.H{"types": types})
}
