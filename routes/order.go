package routes

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/handlers"
)

func OrderRoutes(app *fiber.App, orderHandler *handlers.OrderHandler) {
	app.Get("/orders", orderHandler.GetOrders)
	app.Get("/orders/:order_id", orderHandler.GetOrder)
	app.Post("/orders", orderHandler.CreateOrder)
	app.Patch("/orders/:order_id", orderHandler.UpdateOrder)
}
