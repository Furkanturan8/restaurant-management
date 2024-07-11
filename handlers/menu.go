package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/services"
)

type MenuHandler struct {
	Service *services.MenuService
}

func NewMenuHandler(service *services.MenuService) *MenuHandler {
	return &MenuHandler{
		Service: service,
	}
}

func (mh *MenuHandler) GetMenus(c *fiber.Ctx) error {
	// Handle GET /menus
	return c.SendString("GetMenus endpoint")
}

func (mh *MenuHandler) GetMenu(c *fiber.Ctx) error {
	// Handle GET /menus/:menu_id
	menuID := c.Params("menu_id")
	return c.SendString("GetMenu endpoint, menu ID: " + menuID)
}

func (mh *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	// Handle POST /menus
	return c.SendString("CreateMenu endpoint")
}

func (mh *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	// Handle PATCH /menus/:menu_id
	menuID := c.Params("menu_id")
	return c.SendString("UpdateMenu endpoint, menu ID: " + menuID)
}
