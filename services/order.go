package services

import (
	"gorm.io/gorm"
	model "restaurant-management/models"
)

type OrderService struct {
	DB *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		DB: db,
	}
}

func (os *OrderService) GetOrders(offset, limit int) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64

	if err := os.DB.Model(&model.Order{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := os.DB.Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (os *OrderService) GetOrderByID(id int) (model.Order, error) {
	var order model.Order
	err := os.DB.First(&order, id).Error

	return order, err
}

func (os *OrderService) CreateOrder(order model.Order) error {
	return os.DB.Create(&order).Error
}

func (os *OrderService) UpdateOrder(order model.Order) error {
	return os.DB.Model(&order).Omit("create_at").Updates(order).Error
}
