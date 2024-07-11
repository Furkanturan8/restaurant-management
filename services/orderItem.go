package services

import "database/sql"

type OrderItemService struct {
	DB *sql.DB
}

func NewOrderItemService(db *sql.DB) *OrderItemService {
	return &OrderItemService{DB: db}
}

// orderItem db işlemleri eklenecek
