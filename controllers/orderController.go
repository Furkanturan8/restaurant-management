package controllers

import "github.com/gofiber/fiber/v2"

func GetOrders(c *fiber.Ctx) error {
	// Handle GET /orders
	return c.SendString("GetOrders endpoint")
}

func GetOrder(c *fiber.Ctx) error {
	// Handle GET /orders/:order_id
	orderID := c.Params("order_id")
	return c.SendString("GetOrder endpoint, order ID: " + orderID)
}

func CreateOrder(c *fiber.Ctx) error {
	// Handle POST /orders
	return c.SendString("CreateOrder endpoint")
}

func UpdateOrder(c *fiber.Ctx) error {
	// Handle PATCH /orders/:order_id
	orderID := c.Params("order_id")
	return c.SendString("UpdateOrder endpoint, order ID: " + orderID)
}
