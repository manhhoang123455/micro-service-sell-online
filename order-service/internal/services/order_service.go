package services

import (
	"order-service/internal/models"
	"order-service/internal/repositories"
	"order-service/internal/request"
	"time"
)

type OrderService struct {
	OrderRepo *repositories.OrderRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(input *request.OrderRequest) error {
	order := &models.Order{}
	order.Items = make([]models.OrderItem, len(input.Items))
	err := s.OrderRepo.Create(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) ListOrders(page, limit int, userID uint, startDate, endDate time.Time) ([]models.Order, int64, error) {
	return s.OrderRepo.ListOrders(page, limit, userID, startDate, endDate)
}

func (s *OrderService) GetOrderDetail(id uint) (*models.Order, error) {
	return s.OrderRepo.GetDetailOrder(id)
}
