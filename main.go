package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	database "restaurant-management/database"
	"restaurant-management/routes"
)

func main() {
	fmt.Println("\n--------------------------------BİSMİLLAH--------------------------------\n")
	app := fiber.New()

	// Logger middleware
	app.Use(logger.New())

	// DB Init
	_, err := database.DBInstance()
	if err != nil {
		fmt.Println(err)
	}

	// Burada örneğin spesifik bir tabloya ulaşmak istediğimizde OpenTable fonksiyonuna param. gireceğiz.
	/*usersDB := db.OpenTable("users")
	err = usersDB.Ping()
	if err != nil {
		return
	}*/

	// Define routes
	routes.UserRoutes(app)

	// Authentication middleware
	// app.Use(middleware.Authentication())

	routes.FoodRoutes(app)
	routes.MenuRoutes(app)
	routes.TableRoutes(app)
	routes.OrderRoutes(app)
	routes.OrderItemRoutes(app)
	routes.InvoiceRoutes(app)

	fmt.Println("\n-------------------------ELHAMDÜLİLLAH SORUN YOK-------------------------\n")

	port := "3000"
	err = app.Listen(":" + port)
}
