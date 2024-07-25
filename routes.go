package main

import (
	"github.com/diegopaniago/go-currency-converter/currency/api"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Send([]byte("I am alive!"))
	})

	app.Get("/currency/:code/:target", api.GetCurrency)
}
