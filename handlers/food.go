package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restaurant-management/services"
)

type FoodHandler struct {
	Service *services.FoodService
}

func NewFoodHandler(service *services.FoodService) *FoodHandler {
	return &FoodHandler{Service: service}
}

func (fh *FoodHandler) GetFoods(c *fiber.Ctx) error {

	return nil
}

func (fh *FoodHandler) GetFood(c *fiber.Ctx) error {
	//foodID := c.Params("food_id")

	// FoodService üzerinden foodID'ye göre yiyecek öğesini al
	//food, err := fh.Service.GetFoodByID(foodID)
	//if err != nil {
	//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	//	}
	return c.SendString("GETFood endpoint")
	//	return c.JSON(fiber.Map{"food": food})
}

func (fh *FoodHandler) CreateFood(c *fiber.Ctx) error {
	// Yeni yiyecek öğesi oluşturmak için gerekli işlemleri yap
	// Örneğin, c.Body() üzerinden gelen JSON verisini parse ederek FoodService.CreateFood() ile kaydet
	return c.SendString("CreateFood endpoint")
}

func (fh *FoodHandler) UpdateFood(c *fiber.Ctx) error {
	foodID := c.Params("food_id")

	// Yiyecek öğesini güncellemek için gerekli işlemleri yap
	// Örneğin, c.Body() üzerinden gelen JSON verisini parse ederek FoodService.UpdateFood(foodID) ile güncelle
	return c.SendString("UpdateFood endpoint, food ID: " + foodID)
}
