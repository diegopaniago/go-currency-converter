package settings

import (
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	CurrencyApiUrl string
}

func Load() Envs {
	godotenv.Load()
	currencyApiUrl := os.Getenv("CURRENCY_API_URL")
	if currencyApiUrl == "" {
		currencyApiUrl = "https://test.random"
	}
	return Envs{
		CurrencyApiUrl: currencyApiUrl,
	}
}
