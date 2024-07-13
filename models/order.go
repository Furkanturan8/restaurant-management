package models

import "time"

type Order struct {
	OrderID   uint      `json:"order_id" gorm:"primaryKey;autoIncrement"`
	OrderDate time.Time `json:"order_date" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	TableID   uint      `json:"table_id" validate:"required"`
}
