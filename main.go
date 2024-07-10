package restaurant_management

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"restaurant-management/middleware"
	"restaurant-management/routes"
)

func main() {
	app := fiber.New()

	// Logger middleware
	app.Use(logger.New())

	// Define routes
	routes.UserRoutes(app)

	// Authentication middleware
	app.Use(middleware.Authentication())

	routes.FoodRoutes(app)
	routes.MenuRoutes(app)
	routes.TableRoutes(app)
	routes.OrderRoutes(app)
	routes.OrderItemRoutes(app)
	routes.InvoiceRoutes(app)

	// Start the server
	port := "3000" // Port numarasını buraya ekleyin
	app.Listen(":" + port)

}
