package routes

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/handlers"
)

func TableRoutes(app *fiber.App, tableHandler *handlers.TableHandler) {
	app.Get("/tables", tableHandler.GetTables)
	app.Get("/tables/:table_id", tableHandler.GetTable)
	app.Post("/tables", tableHandler.CreateTable)
	app.Patch("/tables/:table_id", tableHandler.UpdateTable)
}
