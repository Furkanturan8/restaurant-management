package services

import (
	"gorm.io/gorm"
	model "restaurant-management/models"
)

type OrderItemService struct {
	DB *gorm.DB
}

func NewOrderItemService(db *gorm.DB) *OrderItemService {
	return &OrderItemService{DB: db}
}

func (ois *OrderItemService) GetOrderItems(offset, limit int) ([]model.OrderItem, int64, error) {
	var orderItems []model.OrderItem
	var total int64

	if err := ois.DB.Model(&model.OrderItem{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := ois.DB.Offset(offset).Limit(limit).Find(&orderItems).Error; err != nil {
		return nil, 0, err
	}

	return orderItems, total, nil

}

func (ois *OrderItemService) GetOrderItemsByOrder(orderID string) ([]model.OrderItem, error) {
	var orderItems []model.OrderItem

	if err := ois.DB.Where("order_id=?", orderID).Find(&orderItems).Error; err != nil {
		return nil, err
	}

	return orderItems, nil
}

func (ois *OrderItemService) GetOrderItemByID(id int) (model.OrderItem, error) {
	var orderItem model.OrderItem
	//err := ois.DB.First(&orderItem, id).Error
	if err := ois.DB.Where("order_item_id", id).Find(&orderItem).Error; err != nil {
		return model.OrderItem{}, err
	}
	return orderItem, nil
}

func (ois *OrderItemService) CreateOrderItem(orderItem model.OrderItem) error {
	return ois.DB.Create(&orderItem).Error
}

func (ois *OrderItemService) UpdateOrderItem(orderItem model.OrderItem) error {
	return ois.DB.Model(&model.OrderItem{}).Where("order_item_id = ?", orderItem.ID).Omit("created_at").Updates(orderItem).Error
}
