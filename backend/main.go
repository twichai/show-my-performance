package main

import (
	"log"

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

	// Create a user
	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		DB.Create(&user)
		return c.JSON(user)
	})

	// Get all users
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []User
		DB.Find(&users)
		return c.JSON(users)
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
