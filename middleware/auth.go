package middleware

import (
    "github.com/gofiber/fiber/v2"
    "strings"
    "log"
    "fiber/database"
    "time"
)

func AuthMiddleware(c *fiber.Ctx) error {
    // Skip auth check for login and register endpoints
    path := c.Path()
    log.Printf("AuthMiddleware: Checking path %s", path)

    // Skip auth for public paths
    if path == "/login" || path == "/register" || 
       strings.HasPrefix(path, "/js/") || 
       strings.HasPrefix(path, "/css/") || 
       strings.HasPrefix(path, "/images/") {
        log.Printf("AuthMiddleware: Skipping auth check for public path")
        return c.Next()
    }

    // Check if user is authenticated
    sessionToken := c.Cookies("session")
    log.Printf("AuthMiddleware: Session cookie value: %s", sessionToken)

    if sessionToken == "" {
        log.Printf("AuthMiddleware: No session cookie found, redirecting to login")
        return c.Redirect("/login")
    }

    // Validate session token in database
    var userId int
    var createdAt time.Time
    err := database.DB.QueryRow("SELECT user_id, created_at FROM sessions WHERE token = ?", sessionToken).Scan(&userId, &createdAt)
    if err != nil {
        log.Printf("AuthMiddleware: Invalid session token: %v", err)
        return c.Redirect("/login")
    }

    // Check if session is expired (24 hours)
    if time.Since(createdAt) > 24*time.Hour {
        log.Printf("AuthMiddleware: Session expired")
        // Delete expired session
        database.DB.Exec("DELETE FROM sessions WHERE token = ?", sessionToken)
        return c.Redirect("/login")
    }

    // Add user info to context
    c.Locals("userId", userId)
    log.Printf("AuthMiddleware: Authentication successful, proceeding with request for user %d", userId)

    return c.Next()
}