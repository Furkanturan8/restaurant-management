package services

import (
	"gorm.io/gorm"
	model "restaurant-management/models"
)

type InvoiceService struct {
	DB *gorm.DB
}

func NewInvoiceService(db *gorm.DB) *InvoiceService {
	return &InvoiceService{DB: db}
}

func (is *InvoiceService) GetInvoices(offset, limit int) ([]model.Invoice, int64, error) {
	var invoices []model.Invoice
	var total int64

	if err := is.DB.Model(&model.Invoice{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := is.DB.Offset(offset).Limit(limit).Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	return invoices, total, nil
}

func (is *InvoiceService) GetInvoiceByID(id int) (model.Invoice, error) {
	var invoice model.Invoice
	if err := is.DB.Where("invoice_id", id).Find(&invoice).Error; err != nil {
		return model.Invoice{}, err
	}

	return invoice, nil
}

func (is *InvoiceService) CreateInvoice(invoice model.Invoice) error {
	return is.DB.Create(&invoice).Error
}

func (is *InvoiceService) UpdateInvoice(invoice model.Invoice) error {
	return is.DB.Model(&model.Invoice{}).Where("invoice_id = ?", invoice.InvoiceID).Omit("created_at").Updates(invoice).Error

}
