package services

import (
	"gorm.io/gorm"
	model "restaurant-management/models"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (us *UserService) GetUsers(offset, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	if err := us.DB.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := us.DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (us *UserService) GetUserByID(id int) (model.User, error) {
	var user model.User
	err := us.DB.First(&user, id).Error

	return user, err
}
