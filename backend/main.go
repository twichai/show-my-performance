package main

import (
	"log"
	"os"
	"show-my-performance/backend/adapters"
	"show-my-performance/backend/core"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// Initialize SQLite database
func initDB() {
	var err error
	// Configure GORM logger
	newLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Writer for logging
		gormLogger.Config{
			SlowThreshold:             time.Second,     // Log queries slower than this threshold
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound errors
			Colorful:                  true,            // Enable color output
		},
	)

	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger, // Use the configured logger
	})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	DB.AutoMigrate(&core.User{}) // AutoMigrate will create the table if it doesn't exist
	DB.AutoMigrate(&core.Post{}) // AutoMigrate will create the table if it doesn't exist
}

func getCurrentUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)
	email := claims["email"].(string)

	core.CurrentUser.ID = uint(userID)
	core.CurrentUser.Email = email
	return c.Next()
}

func main() {
	app := fiber.New()
	initDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("SECRET_KEY")

	// Fiber logger middleware
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n", // Custom log format
		TimeFormat: "2006-01-02 15:04:05",                       // Time format
		TimeZone:   "Local",                                     // Timezone
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // Allow all origins (use specific domains in production)
		AllowMethods: "GET,POST,PUT,DELETE",                         // Allowed HTTP methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", // Allowed headers, including Authorization
	}))

	userRepository := adapters.NewGormUserRepository(DB)
	userService := core.NewUserService(userRepository)
	userHandler := adapters.NewUserHandler(userService)

	app.Post("/signup", userHandler.RegisterUser)
	app.Post("/login", userHandler.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secretKey)},
	}))
	app.Use(getCurrentUser)

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
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			return c.SendString("Welcome " + claims["email"].(string))
		})
	// Start server
	log.Fatal(app.Listen(":3000"))
}
