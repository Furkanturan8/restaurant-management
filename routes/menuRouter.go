package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "restaurant-management/controllers"
)

func MenuRoutes(app *fiber.App) {
	app.Get("/menus", controller.GetMenus)
	app.Get("/menus/:menu_id", controller.GetMenu)
	app.Post("/menus", controller.CreateMenu)
	app.Patch("/menus/:menu_id", controller.UpdateMenu)
}
