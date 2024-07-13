package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	helper "restaurant-management/helpers"
	"restaurant-management/services"
	"strconv"
)

type UserHandler struct {
	Service   *services.UserService
	Validator *validator.Validate
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service, Validator: validator.New()}
}

func (uh *UserHandler) GetUsers(c *fiber.Ctx) error {
	startIndex, recordPerPage := helper.Pagination(c)

	users, total, err := uh.Service.GetUsers(startIndex, recordPerPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error occurred while listing users"})
	}

	return c.JSON(fiber.Map{
		"total_count": total,
		"users":       users,
	})

}

func (uh *UserHandler) GetUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz user ID"})
	}

	user, err := uh.Service.GetUserByID(int(uint(userID)))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"user": user})
}

func (uh *UserHandler) SignUp(c *fiber.Ctx) error {
	// Handle POST /users/signup
	return c.SendString("SignUp endpoint")
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	// Handle POST /users/login
	return c.SendString("Login endpoint")
}

func HashPassword(password string) string {
	// Implement password hashing
	return password
}

func VerifyPassword(hashedPassword, password string) bool {
	// Implement password verification
	return hashedPassword == password
}
