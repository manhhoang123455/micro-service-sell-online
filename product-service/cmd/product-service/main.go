package main

import (
	"log"
	"product-service/config"
	"product-service/internal/models"
	"product-service/internal/routes"
	"product-service/pkg/database"
	"product-service/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Product Service API
// @version 1.0
// @description This is a product service API.
// @host localhost:8080
// @externalDocs.description  OpenAPI
// @BasePath /
func main() {
	config.LoadConfig()
	database.InitDB()
	utils.InitMinio()
	// Tự động migrate các bảng
	db := database.GetDB()
	if err := db.AutoMigrate(&models.Product{}, &models.Image{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	r := gin.Default()
	routes.RegisterRoutes(r)
	// Thêm route cho Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
