package api

import (
	"encoding/json"
	"strings"

	"github.com/diegopaniago/go-currency-converter/currency/service"
	"github.com/diegopaniago/go-currency-converter/settings"
	"github.com/gofiber/fiber/v3"
)

func GetCurrency(c fiber.Ctx) error {
	currencyProvider := service.CurrencyProvider(service.CurrencyProviderImpl{
		OriginURL: settings.Load().CurrencyApiUrl,
	})
	code := c.Params("code")
	targets := c.Query("targets")
	currency, err := currencyProvider.GetCurrency(code, strings.Split(targets, ","))
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	response, err := json.MarshalIndent(currency, "", "  ")
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	return c.Send([]byte(response))
}
