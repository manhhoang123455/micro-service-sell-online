package routes

import (
	"github.com/gin-gonic/gin"
	"product-service/internal/controllers"
	"product-service/internal/repositories"
	"product-service/internal/services"
	"product-service/pkg/database"
)

func RegisterRoutes(r *gin.Engine) {
	db := database.GetDB()
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)
	productRoutes := r.Group("/products")
	{
		productRoutes.POST("", productController.CreateProduct)
		productRoutes.GET("", productController.GetProductByID)
		productRoutes.GET("/:id", productController.GetProductByID)
		productRoutes.PUT("/:id", productController.UpdateProduct)
	}
}
