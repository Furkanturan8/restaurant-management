package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"restaurant-management/models"
	"restaurant-management/services"
	"strconv"
)

type FoodHandler struct {
	Service *services.FoodService
}

func NewFoodHandler(service *services.FoodService) *FoodHandler {
	return &FoodHandler{Service: service}
}

func (fh *FoodHandler) GetFoods(c *fiber.Ctx) error {
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage", "10"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage

	foods, total, err := fh.Service.GetFoods(startIndex, recordPerPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error occurred while listing food items"})
	}

	return c.JSON(fiber.Map{
		"total_count": total,
		"food_items":  foods,
	})
}

func (fh *FoodHandler) GetFood(c *fiber.Ctx) error {
	foodID, err := strconv.Atoi(c.Params("food_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz food ID"})
	}

	food, err := fh.Service.GetFoodByID(int(uint(foodID)))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"food": food})
}

func (fh *FoodHandler) CreateFood(c *fiber.Ctx) error {

	var food models.Food
	if err := c.BodyParser(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON formatı"})
	}

	err := fh.Service.CreateFood(food)
	if err != nil {
		log.Printf("Food oluşturulurken hata oluştu: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Food oluşturulurken hata oluştu"})
	}

	return c.SendString("Yeni food oluşturuldu")
}

func (fh *FoodHandler) UpdateFood(c *fiber.Ctx) error {

	foodID, err := strconv.Atoi(c.Params("food_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz food ID"})
	}

	var food models.Food
	if err := c.BodyParser(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	food.ID = uint(foodID)
	err = fh.Service.UpdateFood(food)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Food başarılı şekilde güncellendi!"})
}
