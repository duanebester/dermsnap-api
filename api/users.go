package api

import (
	"dermsnap/models"

	"github.com/gofiber/fiber/v2"
)

func (a API) Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	if user == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "user not found"},
		)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
}
