package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	initDB()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
