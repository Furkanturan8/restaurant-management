package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "restaurant-management/controllers"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", controller.GetUsers)
	app.Get("/users/:user_id", controller.GetUser)
	app.Post("/users/signup", controller.SignUp)
	app.Post("/users/login", controller.Login)
}
