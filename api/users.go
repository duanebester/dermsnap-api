package api

import (
	"dermsnap/api/http"
	"dermsnap/models"

	"github.com/gofiber/fiber/v2"
)

func (a API) Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
}

func (a API) GetUserInfo(c *fiber.Ctx, userId http.UserId) error {
	user := c.Locals("user").(*models.User)
	if user.Role != models.Admin && user.ID.String() != userId.String() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	userInfo, err := a.services.UserService.GetUserInfo(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(userInfo)
}

func (a API) CreateUserInfo(c *fiber.Ctx, userId http.UserId) error {
	user := c.Locals("user").(*models.User)
	if user.Role != models.Admin && user.ID.String() != userId.String() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	var createUserInfo models.CreateUserInfo
	if err := c.BodyParser(&createUserInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	userInfo, err := a.services.UserService.CreateUserInfo(userId, createUserInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(userInfo)
}
