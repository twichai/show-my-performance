package main

import (
	"log"
	"show-my-performance/backend/adapters"
	"show-my-performance/backend/core"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app.Use(logger.New(logger.Config{}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                            // Allow all origins (use specific domains in production)
		AllowMethods: "GET,POST,PUT,DELETE",          // Allowed HTTP methods
		AllowHeaders: "Origin, Content-Type, Accept", // Allowed headers
	}))

	userRepository := adapters.NewGormUserRepository(DB)
	userService := core.NewOrderService(userRepository)
	userHandler := adapters.NewUserHandler(userService)

	app.Post("/signup", userHandler.RegisterUser)
	app.Post("/login", userHandler.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	postRepository := adapters.NewGormPostRepository(DB)
	postService := core.NewPostService(postRepository)
	postHandler := adapters.NewPostHandler(postService)

	app.Get("/posts", postHandler.GetAllPosts)
	app.Get("/posts/:id", postHandler.GetPostByID)
	app.Post("/posts", postHandler.CreatePost)
	app.Put("/posts/:id", postHandler.UpdatePost)
	app.Delete("/posts/:id", postHandler.DeletePost)
	app.Get("/posts/user/:userID", postHandler.GetPostsByUserID)

	app.Get("/test",
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "Hello, World!",
			})
		})
	// Start server
	log.Fatal(app.Listen(":3000"))
}
