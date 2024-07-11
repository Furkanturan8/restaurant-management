package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/services"
)

type TableHandler struct {
	Service *services.TableService
}

func NewTableHandler(service *services.TableService) *TableHandler {
	return &TableHandler{Service: service}
}

func (th *TableHandler) GetTables(c *fiber.Ctx) error {
	// Handle GET /tables
	return c.SendString("GetTables endpoint")
}

func (th *TableHandler) GetTable(c *fiber.Ctx) error {
	// Handle GET /tables/:table_id
	tableID := c.Params("table_id")
	return c.SendString("GetTable endpoint, table ID: " + tableID)
}

func (th *TableHandler) CreateTable(c *fiber.Ctx) error {
	// Handle POST /tables
	return c.SendString("CreateTable endpoint")
}

func (th *TableHandler) UpdateTable(c *fiber.Ctx) error {
	// Handle PATCH /tables/:table_id
	tableID := c.Params("table_id")
	return c.SendString("UpdateTable endpoint, table ID: " + tableID)
}
