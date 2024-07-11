package models

import "time"

type Food struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      *string   `json:"name" validate:"required,min=2,max=100"`
	Price     *float64  `json:"price" validate:"required"`
	FoodImage *string   `json:"food_image" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	FoodID    string    `json:"food_id" gorm:"unique"`
	MenuID    *string   `json:"menu_id" validate:"required"`
}
