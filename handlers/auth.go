package handlers

import (
	"database/sql"
	"fiber/database"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		log.Printf("Login error: Failed to parse body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	log.Printf("Login attempt for user: %s", input.Username)

	// Get user from database
	var storedPassword string
	var userId int
	err := database.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", input.Username).Scan(&userId, &storedPassword)
	if err == sql.ErrNoRows {
		log.Printf("Login failed: User not found: %s", input.Username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid credentials",
		})
	} else if err != nil {
		log.Printf("Login error: Database error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Internal server error",
		})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(input.Password)); err != nil {
		log.Printf("Login failed: Invalid password for user: %s", input.Username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid credentials",
		})
	}

	// Generate session token
	sessionToken := uuid.New().String()
	log.Printf("Generated session token for user %s", input.Username)

	// Store session in database
	_, err = database.DB.Exec("INSERT INTO sessions (user_id, token, created_at) VALUES (?, ?, ?)",
		userId, sessionToken, time.Now())
	if err != nil {
		log.Printf("Login error: Failed to store session: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to create session",
		})
	}

	// Set session cookie
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
	})

	log.Printf("Login successful for user: %s", input.Username)
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
	})
}

func Register(c *fiber.Ctx) error {
	var creds Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", creds.Username, hashedPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func Logout(c *fiber.Ctx) error {
	// Get session token
	sessionToken := c.Cookies("session")
	if sessionToken != "" {
		// Delete session from database
		_, err := database.DB.Exec("DELETE FROM sessions WHERE token = ?", sessionToken)
		if err != nil {
			log.Printf("Logout error: Failed to delete session: %v", err)
		}
	}

	// Clear session cookie
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
	})

	log.Printf("User logged out successfully")
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Logged out successfully",
	})
}
