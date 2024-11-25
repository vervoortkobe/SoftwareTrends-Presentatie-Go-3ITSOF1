package handlers

import (
	"fiber/database"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
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
}
