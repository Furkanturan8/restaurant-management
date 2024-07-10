package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "restaurant-management/controllers"
)

func OrderItemRoutes(app *fiber.App) {
	app.Get("/orderItems", controller.GetOrderItems)
	app.Get("/orderItems/:orderItem_id", controller.GetOrderItem)
	app.Get("/orderItems-order/:order_id", controller.GetOrderItemsByOrder)
	app.Post("/orderItems", controller.CreateOrderItem)
	app.Patch("/orderItems/:orderItem_id", controller.UpdateOrderItem)
}
