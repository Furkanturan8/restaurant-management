package models

import "time"

type Menu struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" validate:"required"`
	Category  string     `json:"category" validate:"required"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	MenuID    string     `json:"menu_id" gorm:"unique"`
}
