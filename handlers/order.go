package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/services"
)

type OrderHandler struct {
	Service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

func (oh *OrderHandler) GetOrders(c *fiber.Ctx) error {
	// Handle GET /orders
	return c.SendString("GetOrders endpoint")
}

func (oh *OrderHandler) GetOrder(c *fiber.Ctx) error {
	// Handle GET /orders/:order_id
	orderID := c.Params("order_id")
	return c.SendString("GetOrder endpoint, order ID: " + orderID)
}

func (oh *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	// Handle POST /orders
	return c.SendString("CreateOrder endpoint")
}

func (oh *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	// Handle PATCH /orders/:order_id
	orderID := c.Params("order_id")
	return c.SendString("UpdateOrder endpoint, order ID: " + orderID)
}
