package main

import (
	"github.com/diegopaniago/go-currency-converter/product/api"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Send([]byte("I am alive!"))
	})

	// Product Handlers
	app.Post("/product", api.AddProduct)
	app.Get("/product/:id", api.GetProduct)
}
