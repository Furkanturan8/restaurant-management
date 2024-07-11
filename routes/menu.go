package routes

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/handlers"
)

func MenuRoutes(app *fiber.App, menuHandler *handlers.MenuHandler) {
	app.Get("/menus", menuHandler.GetMenus)
	app.Get("/menus/:menu_id", menuHandler.GetMenu)
	app.Post("/menus", menuHandler.CreateMenu)
	app.Patch("/menus/:menu_id", menuHandler.UpdateMenu)
}
