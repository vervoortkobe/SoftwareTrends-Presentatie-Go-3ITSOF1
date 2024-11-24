package main

import (
	"fiber/handlers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/register", showRegisterPage)
	app.Post("/register", handlers.Register)
	app.Get("/login", showLoginPage)
	app.Post("/login", handlers.Login)
	app.Get("/", showMainPage)
	app.Post("/posts", handlers.CreatePost)
	app.Put("/posts/:id", handlers.EditPost)
	app.Delete("/posts/:id", handlers.DeletePost)
}

func showRegisterPage(c *fiber.Ctx) error {
	return c.SendFile("views/register.html")
}

func showLoginPage(c *fiber.Ctx) error {
	return c.SendFile("views/login.html")
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

	return c.Render("main.html", fiber.Map{
		"Posts": posts,
		"User":  user,
	})
}
