package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "restaurant-management/controllers"
)

func FoodRoutes(app *fiber.App) {
	app.Get("/foods", controller.GetFoods)
	app.Get("/foods/:food_id", controller.GetFood)
	app.Post("/foods", controller.CreateFood)
	app.Patch("/foods/:food_id", controller.UpdateFood)
}
