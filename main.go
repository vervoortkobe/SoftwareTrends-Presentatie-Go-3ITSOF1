package main

import (
	"fiber/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB()
	app := fiber.New()

	setupRoutes(app)

	log.Printf("⚡ | WebServer listening on [http://localhost%s]!\n", ":3000")
	log.Fatal(app.Listen(":3000"))
}
