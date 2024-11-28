package middleware

import (
	"fiber/database"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Skip auth check for login and register endpoints
	path := c.Path()
	fmt.Printf("\nAuthMiddleware: Checking path %s", path)

	// Skip auth for public paths
	if path == "/login" || path == "/register" ||
		strings.HasPrefix(path, "/js/") ||
		strings.HasPrefix(path, "/css/") ||
		strings.HasPrefix(path, "/images/") {
		fmt.Printf("\nAuthMiddleware: Skipping auth check for public path")
		return c.Next()
	}

	// Check if user is authenticated
	sessionToken := c.Cookies("session")
	fmt.Printf("\nAuthMiddleware: Session cookie value: %s", sessionToken)

	if sessionToken == "" {
		fmt.Printf("\nAuthMiddleware: No session cookie found, redirecting to login")
		return c.Redirect("/login")
	}

	// Validate session token in database
	var userId int
	var createdAt time.Time
	err := database.DB.QueryRow("SELECT user_id, created_at FROM sessions WHERE token = ?", sessionToken).Scan(&userId, &createdAt)
	if err != nil {
		fmt.Printf("\nAuthMiddleware: Invalid session token: %v", err)
		return c.Redirect("/login")
	}

	// Check if session is expired (24 hours)
	if time.Since(createdAt) > 24*time.Hour {
		fmt.Printf("\nAuthMiddleware: Session expired")
		// Delete expired session
		database.DB.Exec("DELETE FROM sessions WHERE token = ?", sessionToken)
		return c.Redirect("/login")
	}

	// Add user info to context
	c.Locals("userId", userId)
	fmt.Printf("\nAuthMiddleware: Authentication successful, proceeding with request for user %d", userId)

	return c.Next()
}
