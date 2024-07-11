package services

import "database/sql"

type InvoiceService struct {
	DB *sql.DB
}

func NewInvoiceService(db *sql.DB) *InvoiceService {
	return &InvoiceService{DB: db}
}

// invoice db işlemleri eklenecek
