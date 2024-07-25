package main

import (
	"log"

	"github.com/diegopaniago/go-currency-converter/settings"
	"github.com/gofiber/fiber/v3"
)

func main() {
	settings.Load()
	app := fiber.New()

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
