package controllers

import "github.com/gofiber/fiber/v2"

func GetFoods(c *fiber.Ctx) error {
	// Handle GET /foods
	return c.SendString("GetFoods endpoint")
}

func GetFood(c *fiber.Ctx) error {
	// Handle GET /foods/:food_id
	foodID := c.Params("food_id")
	return c.SendString("GetFood endpoint, food ID: " + foodID)
}

func CreateFood(c *fiber.Ctx) error {
	// Handle POST /foods
	return c.SendString("CreateFood endpoint")
}

func UpdateFood(c *fiber.Ctx) error {
	// Handle PATCH /foods/:food_id
	foodID := c.Params("food_id")
	return c.SendString("UpdateFood endpoint, food ID: " + foodID)
}

func round(num float64) int {
	return 0
}

func toFixed(num float64, precision int) float64 {
	return 0
}
