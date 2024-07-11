package services

import "database/sql"

type MenuService struct {
	DB *sql.DB
}

func NewMenuService(db *sql.DB) *MenuService {
	return &MenuService{DB: db}
}
