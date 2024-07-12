package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	model "restaurant-management/models"
	"restaurant-management/services"
	"strconv"
	"time"
)

type InvoiceHandler struct {
	Service   *services.InvoiceService
	Validator *validator.Validate
}

func NewInvoiceHandler(service *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{Service: service, Validator: validator.New()}
}

func (ih *InvoiceHandler) GetInvoices(c *fiber.Ctx) error {
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage", "10"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage

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

	// Set created_at and updated_at
	now := time.Now()
	invoice.CreatedAt = now
	invoice.UpdatedAt = now

	err := ih.Service.CreateInvoice(invoice)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invoice oluşturulurken hata oluştu!"})
	}

	return c.SendString("Yeni invoice oluşturuldu!")
}

// todo : update ve create ederken boş değer olsa bile sorun çıkarmıyor, bunu düzelt!

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
	validate := validator.New()
	err = validate.Struct(invoice)
	if err != nil {
		log.Fatal("Validation error:", err)
	}

	invoice.InvoiceID = strconv.Itoa(invoiceID)
	invoice.UpdatedAt = time.Now()

	err = ih.Service.UpdateInvoice(invoice)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message:": "Invoice başarıyla güncellendi!"})

}
