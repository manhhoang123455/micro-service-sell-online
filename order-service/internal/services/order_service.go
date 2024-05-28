package services

import (
	"order-service/internal/models"
	"order-service/internal/repositories"
)

type OrderService struct {
	OrderRepo *repositories.OrderRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.OrderRepo.Create(order)
}

func (s *OrderService) ListOrders(page, limit int) ([]models.Order, int64, error) {
	return s.OrderRepo.ListOrders(page, limit)
}

func (s *OrderService) GetOrderDetail(id uint) (*models.Order, error) {
	return s.OrderRepo.GetDetailOrder(id)
}
