package services

import (
	"gorm.io/gorm"
	model "restaurant-management/models"
)

type FoodService struct {
	DB *gorm.DB
}

func NewFoodService(db *gorm.DB) *FoodService {
	return &FoodService{DB: db}
}

// food db işlemleri yapılacak

func (fs *FoodService) GetFoods(offset, limit int) ([]model.Food, int64, error) {
	var foods []model.Food
	var total int64 // toplam food sayısı
	err := fs.DB.Model(&model.Food{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = fs.DB.Offset(offset).Limit(limit).Find(&foods).Error //Limit metodu, çekilecek maksimum kayıt sayısını; Offset metodu, veri çekme işleminin başlangıç noktasını belirler

	return foods, total, err
}

func (fs *FoodService) GetFoodByID(id int) (model.Food, error) {
	var food model.Food
	err := fs.DB.First(&food, id).Error

	return food, err
}

func (fs *FoodService) CreateFood(food model.Food) error {
	return fs.DB.Create(&food).Error
}

func (fs *FoodService) UpdateFood(food model.Food) error {
	return fs.DB.Model(&food).Omit("created_at").Updates(food).Error
}
