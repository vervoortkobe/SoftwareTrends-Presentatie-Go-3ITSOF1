package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"fiber/database"
	"fiber/middleware"
	"fiber/handlers"
	"time"
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
			var userId int
			var createdAt time.Time
			err := database.DB.QueryRow("SELECT user_id, created_at FROM sessions WHERE token = ?", sessionToken).Scan(&userId, &createdAt)
			if err == nil && time.Since(createdAt) < 24*time.Hour {
				log.Printf("User already authenticated, redirecting to home")
				return c.Redirect("/")
			}
			// If session is invalid or expired, clear the cookie
			c.Cookie(&fiber.Cookie{
				Name:     "session",
				Value:    "",
				Expires:  time.Now().Add(-time.Hour),
				HTTPOnly: true,
				Secure:   true,
				SameSite: "Strict",
				Path:     "/",
			})
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

	// Add endpoint to get user data
	app.Get("/api/user", func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(int)
		var username string
		err := database.DB.QueryRow("SELECT username FROM users WHERE id = ?", userId).Scan(&username)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to get user data",
			})
		}
		return c.JSON(fiber.Map{
			"username": username,
		})
	})

	// Add endpoint to get posts
	app.Get("/api/posts", func(c *fiber.Ctx) error {
		currentUserId := c.Locals("userId").(int)
		
		rows, err := database.DB.Query(`
			SELECT p.id, p.title, p.content, p.user_id, u.username 
			FROM posts p 
			JOIN users u ON p.user_id = u.id 
			ORDER BY p.id DESC
		`)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to fetch posts",
			})
		}
		defer rows.Close()

		var posts []fiber.Map
		for rows.Next() {
			var id, userId int
			var title, content, username string
			if err := rows.Scan(&id, &title, &content, &userId, &username); err != nil {
				continue
			}
			posts = append(posts, fiber.Map{
				"id": id,
				"title": title,
				"content": content,
				"userId": userId,
				"username": username,
				"isOwner": userId == currentUserId,
			})
		}

		return c.JSON(fiber.Map{
			"posts": posts,
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// This route is now protected
		return c.SendFile("./views/index.html")
	})

	log.Printf("Server starting on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
