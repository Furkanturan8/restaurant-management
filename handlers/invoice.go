package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	helper "restaurant-management/helpers"
	model "restaurant-management/models"
	"restaurant-management/services"
	"strconv"
)

type InvoiceHandler struct {
	Service   *services.InvoiceService
	Validator *validator.Validate
}

func NewInvoiceHandler(service *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{Service: service, Validator: validator.New()}
}

func (ih *InvoiceHandler) GetInvoices(c *fiber.Ctx) error {

	startIndex, recordPerPage := helper.Pagination(c)

	invoices, total, err := ih.Service.GetInvoices(startIndex, recordPerPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error occurred while listing invoices items"})
	}

	return c.JSON(fiber.Map{
		"total_count": total,
		"invoices":    invoices,
	})
}

func (ih *InvoiceHandler) GetInvoice(c *fiber.Ctx) error {

	invoiceID, err := strconv.Atoi(c.Params("invoice_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz invoice id!"})
	}

	invoice, err := ih.Service.GetInvoiceByID(invoiceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if invoice.CreatedAt.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz invoice id!"})
	}

	return c.JSON(fiber.Map{"invoice": invoice})
}

func (ih *InvoiceHandler) CreateInvoice(c *fiber.Ctx) error {
	var invoice model.Invoice
	if err := c.BodyParser(&invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON formatı"})
	}

	// Validation
	if err := ih.Validator.Struct(invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Kontrol: createdAt veya updatedAt alanları girilmişse hata ver
	if !invoice.CreatedAt.IsZero() || !invoice.UpdatedAt.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "createdAt veya updatedAt alanları manuel olarak doldurulamaz"})
	}

	if *invoice.PaymentMethod == "" || *invoice.PaymentStatus == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "PaymentMethod ve PaymentStatus boş bırakılamaz!"})
	}

	err := ih.Service.CreateInvoice(invoice)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invoice oluşturulurken hata oluştu!"})
	}

	return c.SendString("Yeni invoice oluşturuldu!")
}

func (ih *InvoiceHandler) UpdateInvoice(c *fiber.Ctx) error {

	invoiceID, err := strconv.Atoi(c.Params("invoice_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz invoice id"})
	}

	var invoice model.Invoice
	if err := c.BodyParser(&invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validation
	if err := ih.Validator.Struct(invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	invoice.InvoiceID = uint(invoiceID)
	err = ih.Service.UpdateInvoice(invoice)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message:": "Invoice başarıyla güncellendi!"})

}
