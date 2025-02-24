package routers

import (
	"github.com/gin-gonic/gin"
	"product_management/internal/controllers"
)

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(router *gin.Engine, productController *controllers.ProductController) {
	api := router.Group("/api")
	{
		api.GET("/product", productController.GetProducts)
		api.POST("/product", productController.CreateProduct)
		api.GET("/product/:id", productController.GetProduct)
		api.PUT("/product/:id", productController.UpdateProduct)
		api.POST("/product/upload", productController.UploadImage)
		api.DELETE("/product/:id", productController.DeleteProduct)
		api.GET("/product/type", productController.GetProductTypes)
	}
}
