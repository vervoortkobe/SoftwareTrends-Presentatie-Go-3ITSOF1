package main

import (
	"fiber/handlers"
	"fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Static("/", "./public")
	app.Static("/js", "./public/js")
	app.Static("/css", "./public/css")

	app.Get("/", handlers.MainPage)
	app.Get("/login", handlers.LoginPage)
	app.Get("/register", handlers.RegisterPage)
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)

	app.Use(middleware.AuthMiddleware)

	app.Get("/api/user", handlers.GetUser)
	app.Get("/api/posts", handlers.GetPosts)
	app.Post("/api/posts", handlers.CreatePost)
	app.Put("/api/posts/:id", handlers.UpdatePost)
	app.Delete("/api/posts/:id", handlers.DeletePost)

	app.Get("/api/auth/check", handlers.CheckLogin)
}
