package middleware

import (
	"dermsnap/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func EnrichUser(userService services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("auth_token").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)
		user, err := userService.GetUserByID(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
		c.Locals("user", user)
		return c.Next()
	}
}
