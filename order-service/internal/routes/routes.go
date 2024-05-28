package routes

import (
	"github.com/gin-gonic/gin"
	"order-service/internal/controllers"
	"order-service/internal/repositories"
	"order-service/internal/services"
	"order-service/internal/utils"
)

func RegisterRoutes(r *gin.Engine) {
	orderRepo := repositories.NewOrderRepository(utils.GetDB())
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	r.POST("/orders", orderController.CreateOrder)
	r.GET("/orders", orderController.ListOrder)
	r.GET("/orders/:id", orderController.GetOrderDetail)
}
