package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "restaurant-management/controllers"
)

func InvoiceRoutes(app *fiber.App) {
	app.Get("/invoices", controller.GetInvoices)
	app.Get("/invoices/:invoice_id", controller.GetInvoice)
	app.Post("/invoices", controller.CreateInvoice)
	app.Patch("/invoices/:invoice_id", controller.UpdateInvoice)
}
