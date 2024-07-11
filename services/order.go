package services

import "database/sql"

type OrderService struct {
	DB *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		DB: db,
	}
}

// order db işlemleri yazılacak
