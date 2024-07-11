package models

import "time"

type OrderItem struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Quantity    *string   `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	UnitPrice   *float64  `json:"unit_price" validate:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	FoodID      *string   `json:"food_id" validate:"required"`
	OrderItemID string    `json:"order_item_id" gorm:"unique"`
	OrderID     string    `json:"order_id" validate:"required"`
}
