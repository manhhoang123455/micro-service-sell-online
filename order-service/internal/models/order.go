package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint        `json:"user_id"`
	Status string      `json:"status"`
	Items  []OrderItem `json:"items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
