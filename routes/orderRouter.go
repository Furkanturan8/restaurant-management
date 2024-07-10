package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "restaurant-management/controllers"
)

func OrderRoutes(app *fiber.App) {
	app.Get("/orders", controller.GetOrders)
	app.Get("/orders/:order_id", controller.GetOrder)
	app.Post("/orders", controller.CreateOrder)
	app.Patch("/orders/:order_id", controller.UpdateOrder)
}
