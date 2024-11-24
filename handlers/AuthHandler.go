package handlers

import (
	"fiber/database"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Username string `json:"username,omitempty"`
}

func Register(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(AuthResponse{
			Success: false,
			Message: "Error hashing password",
		})
	}

	_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(AuthResponse{
			Success: false,
			Message: "Username already exists",
		})
	}

	log.Printf("User registered: %s\n", username)
	return c.Status(http.StatusCreated).JSON(AuthResponse{
		Success: true,
		Message: "User registered successfully",
		Username: username,
	})
}

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var storedHashedPassword string
	err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedHashedPassword)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(AuthResponse{
			Success: false,
			Message: "Invalid username or password",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(password))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(AuthResponse{
			Success: false,
			Message: "Invalid username or password",
		})
	}

	log.Printf("User logged in: %s\n", username)
	return c.Status(http.StatusOK).JSON(AuthResponse{
		Success: true,
		Message: "Login successful",
		Username: username,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(AuthResponse{
		Success: true,
		Message: "Logged out successfully",
	})
}
