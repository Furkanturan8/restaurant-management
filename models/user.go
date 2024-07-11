package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	FirstName    *string   `json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string   `json:"last_name" validate:"required,min=2,max=100"`
	Password     *string   `json:"password" validate:"required,min=6"`
	Email        *string   `json:"email" validate:"email,required"`
	Avatar       *string   `json:"avatar"`
	Phone        *string   `json:"phone" validate:"required"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID       string    `json:"user_id" gorm:"unique"`
}