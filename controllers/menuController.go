package controllers

import "github.com/gofiber/fiber/v2"

func GetMenus(c *fiber.Ctx) error {
	// Handle GET /menus
	return c.SendString("GetMenus endpoint")
}

func GetMenu(c *fiber.Ctx) error {
	// Handle GET /menus/:menu_id
	menuID := c.Params("menu_id")
	return c.SendString("GetMenu endpoint, menu ID: " + menuID)
}

func CreateMenu(c *fiber.Ctx) error {
	// Handle POST /menus
	return c.SendString("CreateMenu endpoint")
}

func UpdateMenu(c *fiber.Ctx) error {
	// Handle PATCH /menus/:menu_id
	menuID := c.Params("menu_id")
	return c.SendString("UpdateMenu endpoint, menu ID: " + menuID)
}
