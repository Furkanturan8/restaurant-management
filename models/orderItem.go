package models

import "time"

type OrderItem struct {
	OrderItemID uint      `json:"order_item_id" gorm:"primaryKey;autoIncrement"`
	Quantity    *string   `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	UnitPrice   *float64  `json:"unit_price" validate:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	FoodID      uint      `json:"food_id" validate:"required"`
	OrderID     uint      `json:"order_id" validate:"required"`
}
