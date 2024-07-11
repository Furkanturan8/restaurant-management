package services

import (
	"gorm.io/gorm"
)

type InvoiceService struct {
	DB *gorm.DB
}

func NewInvoiceService(db *gorm.DB) *InvoiceService {
	return &InvoiceService{DB: db}
}

// invoice db işlemleri eklenecek
