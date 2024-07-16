package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	helper "restaurant-management/helpers"
	"restaurant-management/models"
	"restaurant-management/services"
	"strconv"
)

type FoodHandler struct {
	Service   *services.FoodService
	Validator *validator.Validate
}

func NewFoodHandler(service *services.FoodService) *FoodHandler {
	return &FoodHandler{Service: service, Validator: validator.New()}
}

func (fh *FoodHandler) GetFoods(c *fiber.Ctx) error {
	startIndex, recordPerPage := helper.Pagination(c)

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

	// Validation
	if err := fh.Validator.Struct(food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Kontrol: createdAt veya updatedAt alanları girilmişse hata ver
	if !food.CreatedAt.IsZero() || !food.UpdatedAt.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "createdAt veya updatedAt alanları manuel olarak doldurulamaz"})
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

	food.FoodID = uint(foodID)
	err = fh.Service.UpdateFood(food)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Food başarılı şekilde güncellendi!"})
}
