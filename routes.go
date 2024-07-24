package main

import (
	"encoding/json"

	"github.com/diegopaniago/go-currency-converter/currency/service"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Send([]byte("I am alive!"))
	})

	app.Get("/currency/:code/:target", func(c fiber.Ctx) error {
		currencyProvider := service.CurrencyProvider(service.CurrencyProviderImpl{})
		code := c.Params("code")
		target := c.Params("target")
		currency, err := currencyProvider.GetCurrency(code, target)
		if err != nil {
			return c.Status(500).Send([]byte(err.Error()))
		}
		response, err := json.MarshalIndent(currency, "", "  ")
		if err != nil {
			return c.Status(500).Send([]byte(err.Error()))
		}
		return c.Send([]byte(response))
	})
}
