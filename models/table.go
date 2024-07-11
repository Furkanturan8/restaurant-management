package models

import "time"

type Table struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	NumberOfGuests *int      `json:"number_of_guests" validate:"required"`
	TableNumber    *int      `json:"table_number" validate:"required"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	TableID        string    `json:"table_id" gorm:"unique"`
}
