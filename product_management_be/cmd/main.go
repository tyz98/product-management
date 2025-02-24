package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"product_management/internal/config"
	"product_management/internal/controllers"
	"product_management/internal/middleware"
	"product_management/internal/repositories"
	"product_management/internal/routers"
	"product_management/internal/services"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database connection
	dsn := config.DBConfigString()
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	defer db.Close()

	// Initialize Product layers
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// Create the Gin router
	router := gin.Default()

	// static files
	router.Static("/static", "/static")
	// middlewares
	router.Use(middleware.CORSMiddleware())
	// Register routes
	routers.RegisterRoutes(router, productController)

	// Start the service
	router.Run(":9000")
}
