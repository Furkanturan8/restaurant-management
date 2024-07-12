package services

import (
	"gorm.io/gorm"
	model "restaurant-management/models"
)

type TableService struct {
	DB *gorm.DB
}

func NewTableService(db *gorm.DB) *TableService {
	return &TableService{DB: db}
}

func (ts *TableService) GetTables(offset, limit int) ([]model.Table, int64, error) {
	var tables []model.Table
	var total int64

	if err := ts.DB.Model(&model.Table{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := ts.DB.Offset(offset).Limit(limit).Find(&tables).Error; err != nil {
		return nil, 0, err
	}

	return tables, total, nil
}

func (ts *TableService) GetTableByID(id int) (model.Table, error) {
	var table model.Table

	if err := ts.DB.Where("table_id", id).Find(&table).Error; err != nil {
		return model.Table{}, err
	}
	return table, nil
}

func (ts *TableService) CreateTable(table model.Table) error {
	return ts.DB.Create(&table).Error
}

func (ts *TableService) UpdateTable(tableId int, table model.Table) error {
	return ts.DB.Model(&model.Table{}).Where("table_id = ?", tableId).Omit("created_at").Updates(table).Error
}
