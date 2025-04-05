package main

import (
	"log"
	"os"
	postAdapter "show-my-performance/backend/adapters/post"
	userAdapter "show-my-performance/backend/adapters/user"
	postCore "show-my-performance/backend/core/post"
	userCore "show-my-performance/backend/core/user"
	"show-my-performance/backend/model"
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
	DB.AutoMigrate(&model.User{}) // AutoMigrate will create the table if it doesn't exist
	DB.AutoMigrate(&model.Post{}) // AutoMigrate will create the table if it doesn't exist
}

func getCurrentUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)
	email := claims["email"].(string)

	model.CurrentUser.ID = uint(userID)
	model.CurrentUser.Email = email
	return c.Next()
}

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024, // 20 MB limit
	})
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

	userRepository := userAdapter.NewGormUserRepository(DB)
	userService := userCore.NewUserService(userRepository)
	userHandler := userAdapter.NewUserHandler(userService)

	app.Post("/signup", userHandler.RegisterUser)
	app.Post("/login", userHandler.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secretKey)},
	}))
	app.Use(getCurrentUser)

	postRepository := postAdapter.NewGormPostRepository(DB)
	postService := postCore.NewPostService(postRepository)
	postHandler := postAdapter.NewPostHandler(postService)

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
