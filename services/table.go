package services

import "database/sql"

type TableService struct {
	DB *sql.DB
}

func NewTableService(db *sql.DB) *TableService {
	return &TableService{DB: db}
}

// db işlemleri like services/food.go
