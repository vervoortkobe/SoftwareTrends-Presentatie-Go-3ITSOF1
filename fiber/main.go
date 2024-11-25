package main

import (
	"fiber/database"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	app := fiber.New()

	setupRoutes(app)

	fmt.Printf("⚡ | WebServer listening on [http://localhost%s]!\n", ":3000")
	log.Fatal(app.Listen(":3000"))
}
