package main

import (
	"log"
	"show-my-performance/backend/adapters"
	"show-my-performance/backend/core"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

var DB *gorm.DB

// Initialize SQLite database
func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	DB.AutoMigrate(&User{}) // AutoMigrate will create the table if it doesn't exist
}

func main() {
	app := fiber.New()
	initDB()

	userRepository := adapters.NewGormUserRepository(DB)
	userService := core.NewOrderService(userRepository)
	userHandler := adapters.NewUserHandler(userService)

	// Create a user
	app.Post("/users", userHandler.RegisterUser)

	// Get all users
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []User
		DB.Find(&users)
		return c.JSON(users)
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
