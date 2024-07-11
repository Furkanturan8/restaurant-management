package models

import "time"

type Invoice struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	InvoiceID      string    `json:"invoice_id" gorm:"unique"`
	OrderID        string    `json:"order_id"`
	PaymentMethod  *string   `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	PaymentStatus  *string   `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
