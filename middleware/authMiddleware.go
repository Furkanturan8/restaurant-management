package middleware

import (
	"github.com/gofiber/fiber/v2"
	helper "restaurant-management/helpers"
)

func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientToken := c.Get("token")
		if clientToken == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No Authorization header provided"})
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		c.Locals("email", claims.Email)
		c.Locals("first_name", claims.FirstName)
		c.Locals("last_name", claims.LastName)
		c.Locals("uid", claims.UserID)

		return c.Next()
	}
}
