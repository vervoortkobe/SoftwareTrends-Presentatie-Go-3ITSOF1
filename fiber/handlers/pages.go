package handlers

import (
	"fiber/database"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoginPage(c *fiber.Ctx) error {
	sessionToken := c.Cookies("session")
	if sessionToken != "" {
		var userId int
		var createdAt time.Time
		err := database.DB.QueryRow("SELECT user_id, created_at FROM sessions WHERE token = ?", sessionToken).Scan(&userId, &createdAt)
		if err == nil && time.Since(createdAt) < 24*time.Hour {
			fmt.Printf("User already authenticated, redirecting to home")
			return c.Redirect("/")
		}
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
}

func RegisterPage(c *fiber.Ctx) error {
	return c.SendFile("./views/register.html")
}

func MainPage(c *fiber.Ctx) error {
	return c.SendFile("./views/index.html")
}
