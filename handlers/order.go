package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	helper "restaurant-management/helpers"
	model "restaurant-management/models"
	"restaurant-management/services"
	"strconv"
)

type OrderHandler struct {
	Service   *services.OrderService
	Validator *validator.Validate
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{Service: service, Validator: validator.New()}
}

func (oh *OrderHandler) GetOrders(c *fiber.Ctx) error {
	startIndex, recordPerPage := helper.Pagination(c)

	orders, total, err := oh.Service.GetOrders(startIndex, recordPerPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Order items listelenirken hata oluştu!"})
	}

	return c.JSON(fiber.Map{"total_count": total, "order_items": orders})

}

func (oh *OrderHandler) GetOrder(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("order_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz Order ID"})
	}

	order, err := oh.Service.GetOrderByID(orderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"order:": order})
}

func (oh *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order model.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON formatı"})
	}

	if err := oh.Validator.Struct(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Kontrol: createdAt veya updatedAt alanları girilmişse hata ver
	if !order.CreatedAt.IsZero() || !order.UpdatedAt.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "createdAt veya updatedAt alanları manuel olarak doldurulamaz"})
	}

	err := oh.Service.CreateOrder(order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Order oluşturulamadı!"})
	}

	return c.SendString("Yeni order oluşturuldu!")
}

func (oh *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("order_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz Order ID"})
	}

	var order model.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	order.OrderID = uint(orderID)
	err = oh.Service.UpdateOrder(order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Order başarıyla güncellendi!"})
}
