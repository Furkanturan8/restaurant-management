package models

import "time"

type Menu struct {
	MenuID    uint       `json:"menu_id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" validate:"required"`
	Category  string     `json:"category" validate:"required"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
