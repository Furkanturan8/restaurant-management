package services

import (
	"database/sql"
	model "restaurant-management/models"
)

type FoodService struct {
	DB *sql.DB
}

func NewFoodService(db *sql.DB) *FoodService {
	return &FoodService{DB: db}
}

// food db işlemleri yapılacak

func (fs *FoodService) GetFoodByID(id int) (model.Food, error) {
	return model.Food{}, nil
}

func (fs *FoodService) CreateFood(food model.Food) error {
	return nil
}
