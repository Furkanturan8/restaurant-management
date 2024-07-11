package routes

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/handlers"
)

func FoodRoutes(app *fiber.App, foodHandler *handlers.FoodHandler) {
	app.Get("/foods", foodHandler.GetFoods)
	app.Get("/foods/:food_id", foodHandler.GetFood)
	app.Post("/foods", foodHandler.CreateFood)
	app.Patch("/foods/:food_id", foodHandler.UpdateFood)
}
