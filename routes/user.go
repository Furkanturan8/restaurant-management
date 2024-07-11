package routes

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/handlers"
)

func UserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	app.Get("/users", userHandler.GetUsers)
	app.Get("/users/:user_id", userHandler.GetUser)
	app.Post("/users/signup", userHandler.SignUp)
	app.Post("/users/login", userHandler.Login)
}
