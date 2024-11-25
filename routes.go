package main

import (
	"fiber/handlers"
	"fiber/middleware"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// Public routes
	app.Get("/register", showRegisterPage)
	app.Post("/register", handlers.Register)
	app.Get("/login", showLoginPage)
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)

	// Auth check endpoint
	app.Get("/api/auth/check", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"user":    c.Locals("user"),
		})
	})

	// Protected routes
	app.Use(middleware.AuthMiddleware)
	app.Get("/", showMainPage)
	app.Post("/posts", handlers.CreatePost)
	app.Put("/posts/:id", handlers.EditPost)
	app.Delete("/posts/:id", handlers.DeletePost)
}

func showLoginPage(c *fiber.Ctx) error {
	return c.SendFile("views/login.html")
}

func showRegisterPage(c *fiber.Ctx) error {
	return c.SendFile("views/register.html")
}

func showMainPage(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Redirect("/login")
	}

	posts, err := handlers.FetchAllPosts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching posts")
	}

	return c.Render("views/index.html", fiber.Map{
		"Posts": posts,
		"User":  user,
	})
}
