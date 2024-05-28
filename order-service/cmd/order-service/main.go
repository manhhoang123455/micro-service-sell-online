package main

import (
	"log"
	"order-service/config"
	"order-service/internal/models"
	"order-service/internal/routes"
	"order-service/internal/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Order Service API
// @version 1.0
// @description This is a order service API.
// @host localhost:8080
// @externalDocs.description  OpenAPI
// @BasePath /
func main() {
	config.LoadConfig()
	utils.InitDB()
	r := gin.Default()
	db := utils.GetDB()
	if err := db.AutoMigrate(&models.Order{}, &models.OrderItem{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
