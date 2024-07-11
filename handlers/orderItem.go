package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/services"
)

type OrderItemHandler struct {
	Service *services.OrderItemService
}

func NewOrderItemHandler(service *services.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{Service: service}
}

func (oih *OrderItemHandler) GetOrderItems(c *fiber.Ctx) error {
	// Handle GET /orderItems logic here
	return c.SendString("GetOrderItems endpoint")
}

func (oih *OrderItemHandler) GetOrderItem(c *fiber.Ctx) error {
	// Handle GET /orderItems/:orderItem_id logic here
	orderItemID := c.Params("orderItem_id")
	return c.SendString("GetOrderItem endpoint, orderItem ID: " + orderItemID)
}

func (oih *OrderItemHandler) GetOrderItemsByOrder(c *fiber.Ctx) error {
	// Handle GET /orderItems-order/:order_id logic here
	orderID := c.Params("order_id")
	return c.SendString("GetOrderItemsByOrder endpoint, order ID: " + orderID)
}

func (oih *OrderItemHandler) CreateOrderItem(c *fiber.Ctx) error {
	// Handle POST /orderItems logic here
	return c.SendString("CreateOrderItem endpoint")
}

func (oih *OrderItemHandler) UpdateOrderItem(c *fiber.Ctx) error {
	// Handle PATCH /orderItems/:orderItem_id logic here
	orderItemID := c.Params("orderItem_id")
	return c.SendString("UpdateOrderItem endpoint, orderItem ID: " + orderItemID)
}

/* Halledicez inşallah
func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	return nil, err
}
*/
