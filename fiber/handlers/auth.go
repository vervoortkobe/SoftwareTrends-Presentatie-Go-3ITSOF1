package handlers

import (
	"database/sql"
	"fiber/database"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckLogin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"user":    c.Locals("user"),
	})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		fmt.Printf("\nLogin error: Failed to parse body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	fmt.Printf("\nLogin attempt for user: %s", input.Username)

	var storedPassword string
	var userId int
	err := database.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", input.Username).Scan(&userId, &storedPassword)
	if err == sql.ErrNoRows {
		fmt.Printf("\nLogin failed: User not found: %s", input.Username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid credentials",
		})
	} else if err != nil {
		fmt.Printf("\nLogin error: Database error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Internal server error",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(input.Password)); err != nil {
		fmt.Printf("\nLogin failed: Invalid password for user: %s", input.Username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid credentials",
		})
	}

	sessionToken := uuid.New().String()
	fmt.Printf("\nGenerated session token for user %s", input.Username)

	_, err = database.DB.Exec("INSERT INTO sessions (user_id, token, created_at) VALUES (?, ?, ?)",
		userId, sessionToken, time.Now())
	if err != nil {
		fmt.Printf("\nLogin error: Failed to store session: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to create session",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
	})

	fmt.Printf("\nLogin successful for user: %s", input.Username)
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

	// Log in the user by creating a session
	var userId int
	err = database.DB.QueryRow("SELECT id FROM users WHERE username = ?", creds.Username).Scan(&userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user ID",
		})
	}

	sessionToken := uuid.New().String()
	_, err = database.DB.Exec("INSERT INTO sessions (user_id, token, created_at) VALUES (?, ?, ?)",
		userId, sessionToken, time.Now())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create session",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func Logout(c *fiber.Ctx) error {
	sessionToken := c.Cookies("session")
	if sessionToken != "" {
		// Delete session from database
		_, err := database.DB.Exec("DELETE FROM sessions WHERE token = ?", sessionToken)
		if err != nil {
			fmt.Printf("\nLogout error: Failed to delete session: %v", err)
		}
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

	fmt.Printf("\nUser logged out successfully")
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Logged out successfully",
	})
}
