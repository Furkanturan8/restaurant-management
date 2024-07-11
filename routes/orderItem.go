package routes

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/handlers"
)

func OrderItemRoutes(app *fiber.App, orderItemHandler *handlers.OrderItemHandler) {
	app.Get("/orderItems", orderItemHandler.GetOrderItems)
	app.Get("/orderItems/:orderItem_id", orderItemHandler.GetOrderItem)
	app.Get("/orderItems-order/:order_id", orderItemHandler.GetOrderItemsByOrder)
	app.Post("/orderItems", orderItemHandler.CreateOrderItem)
	app.Patch("/orderItems/:orderItem_id", orderItemHandler.UpdateOrderItem)
}