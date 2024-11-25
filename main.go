package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"fiber/database"
	"fiber/middleware"
	"fiber/handlers"
)

func main() {
	// Initialize database
	database.InitDB()
	defer database.DB.Close()

	app := fiber.New()

	// Serve static files
	app.Static("/", "./public")
	app.Static("/js", "./public/js")
	app.Static("/css", "./public/css")

	// Public routes (no auth required)
	app.Get("/login", func(c *fiber.Ctx) error {
		// Check if user is already authenticated
		sessionToken := c.Cookies("session")
		if sessionToken != "" {
			// Verify session in database
			var exists bool
			err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE token = ?)", sessionToken).Scan(&exists)
			if err == nil && exists {
				return c.Redirect("/")
			}
		}
		return c.SendFile("./views/login.html")
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.SendFile("./views/register.html")
	})

	// Auth endpoints
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)
	app.Post("/logout", handlers.Logout)

	// Protected routes (auth required)
	app.Use(middleware.AuthMiddleware) // Apply middleware to all routes below this

	app.Get("/", func(c *fiber.Ctx) error {
		// This route is now protected
		return c.SendFile("./views/index.html")
	})

	log.Printf("Server starting on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
