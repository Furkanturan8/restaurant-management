package services

import (
	"gorm.io/gorm"
)

type OrderItemService struct {
	DB *gorm.DB
}

func NewOrderItemService(db *gorm.DB) *OrderItemService {
	return &OrderItemService{DB: db}
}

// orderItem db işlemleri eklenecek
