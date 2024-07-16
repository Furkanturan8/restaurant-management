package middleware

import (
	"github.com/gofiber/fiber/v2"
	helper "restaurant-management/helpers"
	"strings"
)

func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientToken := c.Get("Authorization")
		if clientToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No Authorization header provided"})
		}

		// Token'ı "Bearer " ile birlikte al
		tokenParts := strings.Split(clientToken, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header format"})
		}
		token := tokenParts[1]

		claims, err := helper.ValidateToken(token)
		if err != "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
		}

		c.Locals("email", claims.Email)
		c.Locals("first_name", claims.FirstName)
		c.Locals("last_name", claims.LastName)
		c.Locals("uid", claims.UserID)

		return c.Next()
	}
}
