package controllers

import "github.com/gofiber/fiber/v2"

func GetInvoices(c *fiber.Ctx) error {
	// Handle GET /invoices
	return c.SendString("GetInvoices endpoint")
}

func GetInvoice(c *fiber.Ctx) error {
	// Handle GET /invoices/:invoice_id
	invoiceID := c.Params("invoice_id")
	return c.SendString("GetInvoice endpoint, invoice ID: " + invoiceID)
}

func CreateInvoice(c *fiber.Ctx) error {
	// Handle POST /invoices
	return c.SendString("CreateInvoice endpoint")
}

func UpdateInvoice(c *fiber.Ctx) error {
	// Handle PATCH /invoices/:invoice_id
	invoiceID := c.Params("invoice_id")
	return c.SendString("UpdateInvoice endpoint, invoice ID: " + invoiceID)
}
