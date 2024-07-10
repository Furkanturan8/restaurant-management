package controllers

import "github.com/gofiber/fiber/v2"

func GetTables(c *fiber.Ctx) error {
	// Handle GET /tables
	return c.SendString("GetTables endpoint")
}

func GetTable(c *fiber.Ctx) error {
	// Handle GET /tables/:table_id
	tableID := c.Params("table_id")
	return c.SendString("GetTable endpoint, table ID: " + tableID)
}

func CreateTable(c *fiber.Ctx) error {
	// Handle POST /tables
	return c.SendString("CreateTable endpoint")
}

func UpdateTable(c *fiber.Ctx) error {
	// Handle PATCH /tables/:table_id
	tableID := c.Params("table_id")
	return c.SendString("UpdateTable endpoint, table ID: " + tableID)
}
