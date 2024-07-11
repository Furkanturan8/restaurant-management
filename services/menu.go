package services

import (
	"gorm.io/gorm"
)

type MenuService struct {
	DB *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{DB: db}
}
