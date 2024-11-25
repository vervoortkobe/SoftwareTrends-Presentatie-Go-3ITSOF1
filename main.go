package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"fiber/database"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	app := fiber.New()

	setupRoutes(app)

	log.Printf("Server starting on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}