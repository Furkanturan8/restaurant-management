package routes

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/handlers"
)

func InvoiceRoutes(app *fiber.App, invoiceHandler *handlers.InvoiceHandler) {
	app.Get("/invoices", invoiceHandler.GetInvoices)
	app.Get("/invoices/:invoice_id", invoiceHandler.GetInvoice)
	app.Post("/invoices", invoiceHandler.CreateInvoice)
	app.Patch("/invoices/:invoice_id", invoiceHandler.UpdateInvoice)
}
