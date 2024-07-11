package services

import (
	"gorm.io/gorm"
	model "restaurant-management/models"
)

type MenuService struct {
	DB *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{DB: db}
}

func (ms *MenuService) GetMenus(offset, limit int) ([]model.Menu, int64, error) {
	var menus []model.Menu
	var total int64

	if err := ms.DB.Model(&model.Menu{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := ms.DB.Offset(offset).Limit(limit).Find(&menus).Error; err != nil {
		return nil, 0, err
	}

	return menus, total, nil
}

func (ms *MenuService) GetMenuByID(id int) (model.Menu, error) {
	var menu model.Menu
	err := ms.DB.First(&menu, id).Error

	return menu, err
}

func (ms *MenuService) CreateMenu(menu model.Menu) error {
	return ms.DB.Create(&menu).Error
}

func (ms *MenuService) UpdateMenu(menu model.Menu) error {
	return ms.DB.Model(&menu).Omit("created_at").Updates(menu).Error
}
