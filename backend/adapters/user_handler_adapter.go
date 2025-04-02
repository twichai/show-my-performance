package adapters

import (
	"show-my-performance/backend/core"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService core.UserService
}

func NewUserHandler(userService core.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var user core.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.userService.RegisterUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
