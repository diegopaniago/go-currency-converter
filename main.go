package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
