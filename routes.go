package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/register", showRegisterPage)
	app.Post("/register", register)
	app.Get("/login", showLoginPage)
	app.Post("/login", login)
	app.Get("/", showMainPage)
	app.Post("/posts", createPost)
	app.Put("/posts/:id", editPost)
	app.Delete("/posts/:id", deletePost)
}

func showRegisterPage(c *fiber.Ctx) error {
	return c.SendFile("templates/register.html")
}

func showLoginPage(c *fiber.Ctx) error {
	return c.SendFile("templates/login.html")
}

func showMainPage(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Redirect("/login")
	}

	posts, err := fetchAllPosts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching posts")
	}

	return c.Render("main.html", fiber.Map{
		"Posts": posts,
		"User":  user,
	})
}
