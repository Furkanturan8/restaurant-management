package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "restaurant-management/controllers"
)

func TableRoutes(app *fiber.App) {
	app.Get("/tables", controller.GetTables)
	app.Get("/tables/:table_id", controller.GetTable)
	app.Post("/tables", controller.CreateTable)
	app.Patch("/tables/:table_id", controller.UpdateTable)
}
