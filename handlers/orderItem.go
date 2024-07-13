package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	helper "restaurant-management/helpers"
	"restaurant-management/models"
	"restaurant-management/services"
	"strconv"
)

type OrderItemHandler struct {
	Service   *services.OrderItemService
	Validator *validator.Validate
}

func NewOrderItemHandler(service *services.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{Service: service, Validator: validator.New()}
}

func (oih *OrderItemHandler) GetOrderItems(c *fiber.Ctx) error {
	startIndex, recordPerPage := helper.Pagination(c)

	orderItems, total, err := oih.Service.GetOrderItems(startIndex, recordPerPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Order Items Listelenirken Hata Oluştu!"})
	}

	return c.JSON(fiber.Map{
		"total_count": total,
		"order_items": orderItems,
	})

}

func (oih *OrderItemHandler) GetOrderItem(c *fiber.Ctx) error {
	orderItemID, err := strconv.Atoi(c.Params("order_item_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz order item ID"})
	}

	orderItem, err := oih.Service.GetOrderItemByID(orderItemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"order_item": orderItem})
}

func (oih *OrderItemHandler) GetOrderItemsByOrder(c *fiber.Ctx) error {
	orderID := c.Params("order_id")

	orderItems, err := oih.Service.GetOrderItemsByOrder(orderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Sipariş öğeleri sipariş kimliğine göre listelenirken hata oluştu!"})
	}

	return c.JSON(fiber.Map{"order_items": orderItems})
}

func (oih *OrderItemHandler) CreateOrderItem(c *fiber.Ctx) error {
	var orderItem models.OrderItem
	if err := c.BodyParser(&orderItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON format"})
	}

	if err := oih.Validator.Struct(orderItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Kontrol: createdAt veya updatedAt alanları girilmişse hata ver
	if !orderItem.CreatedAt.IsZero() || !orderItem.UpdatedAt.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "createdAt veya updatedAt alanları manuel olarak doldurulamaz"})
	}

	err := oih.Service.CreateOrderItem(orderItem)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Order item oluşturulurken hata oluştu!"})
	}

	return c.SendString("Yeni order item oluşturuldu!")
}

func (oih *OrderItemHandler) UpdateOrderItem(c *fiber.Ctx) error {
	orderItemID, err := strconv.Atoi(c.Params("order_item_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz order item ID"})
	}

	var orderItem models.OrderItem
	if err := c.BodyParser(&orderItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	orderItem.OrderItemID = uint(orderItemID)
	err = oih.Service.UpdateOrderItem(orderItem)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Order item başarıyla güncellendi!"})

}
