package models

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	OrderDate time.Time `json:"order_date" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	OrderID   string    `json:"order_id" gorm:"unique"`
	TableID   *string   `json:"table_id" validate:"required"`
}
