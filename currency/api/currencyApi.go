package api

import (
	"encoding/json"

	"github.com/diegopaniago/go-currency-converter/currency/service"
	"github.com/diegopaniago/go-currency-converter/settings"
	"github.com/gofiber/fiber/v3"
)

func GetCurrency(c fiber.Ctx) error {
	currencyProvider := service.CurrencyProvider(service.CurrencyProviderImpl{
		OriginURL: settings.Load().CurrencyApiUrl,
	})
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
}
