package userAdapter

import (
	"os"
	userCore "show-my-performance/backend/core/user"
	"show-my-performance/backend/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserHandler struct {
	userService userCore.UserService
}

func NewUserHandler(userService userCore.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		print(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	foundUser, err := h.userService.Login(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password"})
	}

	claims := jwt.MapClaims{
		"user_id": foundUser.ID,
		"email":   foundUser.Email,
		"name":    foundUser.Username,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("SECRET_KEY")
	t, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t, "user": foundUser})

}
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.userService.RegisterUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
