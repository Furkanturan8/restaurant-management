package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/services"
)

type InvoiceHandler struct {
	Service *services.InvoiceService
}

func NewInvoiceHandler(service *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{Service: service}
}

func (ih *InvoiceHandler) GetInvoices(c *fiber.Ctx) error {
	// Handle GET /invoices
	return c.SendString("GetInvoices endpoint")
}

func (ih *InvoiceHandler) GetInvoice(c *fiber.Ctx) error {
	// Handle GET /invoices/:invoice_id
	invoiceID := c.Params("invoice_id")
	return c.SendString("GetInvoice endpoint, invoice ID: " + invoiceID)
}

func (ih *InvoiceHandler) CreateInvoice(c *fiber.Ctx) error {
	// Handle POST /invoices
	return c.SendString("CreateInvoice endpoint")
}

func (ih *InvoiceHandler) UpdateInvoice(c *fiber.Ctx) error {
	// Handle PATCH /invoices/:invoice_id
	invoiceID := c.Params("invoice_id")
	return c.SendString("UpdateInvoice endpoint, invoice ID: " + invoiceID)
}
