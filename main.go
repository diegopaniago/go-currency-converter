package main

import (
	"log"

	"github.com/diegopaniago/go-currency-converter/settings"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	settings.Load()
	app := fiber.New()
	app.Use(cors.New())

	SetupRoutes(app)

	log.Fatal(app.Listen(":5001"))
}
