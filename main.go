package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
	"restaurant-management/database"
	"restaurant-management/handlers"
	"restaurant-management/middleware"
	"restaurant-management/routes"
	"restaurant-management/services"
)

func main() {
	fmt.Println("\n--------------BİSMİLLAH--------------\n")
	app := fiber.New()

	// Logger middleware
	app.Use(logger.New())

	// DB Init
	db, err := database.DBInstance()
	if err != nil {
		fmt.Println(err)
	}

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)
	routes.UserRoutes(app, userHandler)

	// Authentication middleware
	app.Use(middleware.Authentication())

	// Initialize services and handlers
	initializeHandlers(app, db)

	port := "3000"
	err = app.Listen(":" + port)
}

func initializeHandlers(app *fiber.App, db *gorm.DB) {
	// Initialize all modules

	foodService := services.NewFoodService(db)
	foodHandler := handlers.NewFoodHandler(foodService)
	routes.FoodRoutes(app, foodHandler)

	menuService := services.NewMenuService(db)
	menuHandler := handlers.NewMenuHandler(menuService)
	routes.MenuRoutes(app, menuHandler)

	tableService := services.NewTableService(db)
	tableHandler := handlers.NewTableHandler(tableService)
	routes.TableRoutes(app, tableHandler)

	orderService := services.NewOrderService(db)
	orderHandler := handlers.NewOrderHandler(orderService)
	routes.OrderRoutes(app, orderHandler)

	orderItemService := services.NewOrderItemService(db)
	orderItemHandler := handlers.NewOrderItemHandler(orderItemService)
	routes.OrderItemRoutes(app, orderItemHandler)

	invoiceService := services.NewInvoiceService(db)
	invoiceHandler := handlers.NewInvoiceHandler(invoiceService)
	routes.InvoiceRoutes(app, invoiceHandler)
}
