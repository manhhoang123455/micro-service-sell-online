package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/internal/models"
	"order-service/internal/services"
	"order-service/internal/utils"
	"strconv"
)

type OrderController struct {
	OrderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body Order true "Order object"
// @Success 201 {object} Order
// @Router /orders [post]

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := oc.OrderService.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorsResponseByCode(http.StatusInternalServerError, "Failed to create product", utils.SignatureFailed, nil))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Product created successfully", order))
}

func (oc *OrderController) ListOrder(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	orders, total, err := oc.OrderService.ListOrders(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorsResponseByCode(http.StatusInternalServerError, "Failed to list orders", utils.SignatureFailed, nil))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessPageResponse(http.StatusOK, "Orders retrieved successfully", int64(page), int64(limit), total, (total+int64(limit)-1)/int64(limit), orders))
}

func (oc *OrderController) GetOrderDetail(c *gin.Context) {
	orderId, _ := strconv.Atoi(c.Param("id"))
	order, err := oc.OrderService.GetOrderDetail(uint(orderId))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorsResponseByCode(http.StatusBadRequest, "Invalid ID", utils.IdInvalid, nil))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Order retrieved successfully", order))

}
