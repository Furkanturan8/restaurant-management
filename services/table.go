package services

import (
	"gorm.io/gorm"
)

type TableService struct {
	DB *gorm.DB
}

func NewTableService(db *gorm.DB) *TableService {
	return &TableService{DB: db}
}

// db işlemleri like services/food.go
