package services

import (
	"gorm.io/gorm"
	model "restaurant-management/models"
	"time"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (us *UserService) CreateUser(user *model.User) error {
	return us.DB.Create(user).Error
}

func (us *UserService) GetUsers(offset, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	if err := us.DB.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := us.DB.Model(&model.User{}).Omit("password").Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (us *UserService) GetUserByID(id int) (model.User, error) {
	var user model.User
	err := us.DB.First(&user, id).Error

	return user, err
}

func (us *UserService) CountByEmail(email string) (int64, error) {
	var count int64
	err := us.DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count, err
}

func (us *UserService) CountByPhone(phone string) (int64, error) {
	var count int64
	err := us.DB.Model(&model.User{}).Where("phone = ?", phone).Count(&count).Error
	return count, err
}

func (us *UserService) FindByEmail(email string, user *model.User) error {
	return us.DB.Where("email = ?", email).First(user).Error
}

func (us *UserService) UpdateTokens(userID uint, token string, refreshToken string) error {
	return us.DB.Model(&model.User{}).Where("user_id = ?", userID).Updates(map[string]interface{}{
		"token":         token,
		"refresh_token": refreshToken,
		"updated_at":    time.Now(),
	}).Error
}

func (us *UserService) ClearTokens(userID uint) error {
	return us.DB.Model(&model.User{}).Where("user_id = ?", userID).Updates(map[string]interface{}{
		"token":         "",
		"refresh_token": "",
		"updated_at":    time.Now(),
	}).Error
}
