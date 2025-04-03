package main

import (
	"log"
	"show-my-performance/backend/adapters"
	"show-my-performance/backend/core"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	jwtware "github.com/gofiber/contrib/jwt"
)

var DB *gorm.DB

// Initialize SQLite database
func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	DB.AutoMigrate(&core.User{}) // AutoMigrate will create the table if it doesn't exist
}

func main() {
	app := fiber.New()
	initDB()

	userRepository := adapters.NewGormUserRepository(DB)
	userService := core.NewOrderService(userRepository)
	userHandler := adapters.NewUserHandler(userService)

	// Create a user
	app.Post("/users", userHandler.RegisterUser)
	app.Get("/login", userHandler.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Get("/test",
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "Hello, World!",
			})
		})
	// Start server
	log.Fatal(app.Listen(":3000"))
}
