package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	model "restaurant-management/models"
	"restaurant-management/services"
	"strconv"
	"time"
)

type TableHandler struct {
	Service   *services.TableService
	Validator *validator.Validate
}

func NewTableHandler(service *services.TableService) *TableHandler {
	return &TableHandler{Service: service, Validator: validator.New()}
}

func (th *TableHandler) GetTables(c *fiber.Ctx) error {
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage", "10"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	tables, total, err := th.Service.GetTables(startIndex, recordPerPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Tables Listelenirken Hata Oluştu!"})
	}

	return c.JSON(fiber.Map{
		"total_count": total,
		"tables":      tables,
	})
}

func (th *TableHandler) GetTable(c *fiber.Ctx) error {
	tableID, err := strconv.Atoi(c.Params("table_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz table ID"})
	}

	table, err := th.Service.GetTableByID(tableID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"table": table})
}

func (th *TableHandler) CreateTable(c *fiber.Ctx) error {
	var table model.Table
	if err := c.BodyParser(&table); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON format"})
	}

	if err := th.Validator.Struct(table); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	now := time.Now()
	table.CreatedAt = now
	table.UpdatedAt = now

	err := th.Service.CreateTable(table)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Table oluşturulurken hata oluştu!"})
	}

	return c.SendString("Yeni Table Oluşturuldu!")
}

func (th *TableHandler) UpdateTable(c *fiber.Ctx) error {
	tableID, err := strconv.Atoi(c.Params("table_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Id geçersiz!", "tableID": tableID})
	}

	var table model.Table
	if err := c.BodyParser(&table); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	table.UpdatedAt = time.Now()

	err = th.Service.UpdateTable(tableID, table)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Table başarıyla güncellendi!"})
}
