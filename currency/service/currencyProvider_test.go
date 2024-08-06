package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/diegopaniago/go-currency-converter/currency/model"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrency(t *testing.T) {
	tests := []struct {
		name                string
		mockResponse        string
		mockStatusCode      int
		code                string
		targets             []string
		expectError         bool
		expectedResult      []model.Currency
		expectedErrorResult string
	}{
		{
			name:           "Currency Provider should return the currencies",
			mockResponse:   "",
			mockStatusCode: http.StatusOK,
			code:           "USD",
			targets:        []string{"BRL", "EUR"},
			expectError:    false,
			expectedResult: []model.Currency{
				{
					Code:  "USD",
					Name:  "Dollar",
					Price: 1.0000,
					Exchange: model.Exchange{
						Code:  "BRL",
						Name:  "Real",
						Price: 5.15,
					},
				},
				{
					Code:  "USD",
					Name:  "Dollar",
					Price: 1.0000,
					Exchange: model.Exchange{
						Code:  "EUR",
						Name:  "Euro",
						Price: 5.15,
					},
				},
			},
		},
		{
			name:                "Currency Provider should return error if called with invalid code",
			mockResponse:        "",
			mockStatusCode:      http.StatusNotFound,
			code:                "DOLAR",
			targets:             []string{"BRL"},
			expectError:         true,
			expectedErrorResult: "error fetching currency DOLAR to BRL: 404 Not Found",
		},
		{
			name:                "Currency Provider should return error if the external api returns with invalid JSON",
			mockResponse:        `invalid JSON`,
			mockStatusCode:      http.StatusOK,
			code:                "BRL",
			targets:             []string{"USD", "EUR"},
			expectError:         true,
			expectedErrorResult: "error to unmarshal json: invalid character 'i' looking for beginning of value",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/USD-BRL" {
					tc.mockResponse = `[{"code":"USD","codein":"BRL","name":"Dollar/Real","high":"5.20","low":"5.10","varBid":"0.001","pctChange":"0.02","bid":"5.15","ask":"5.16","timestamp":"1609459200","create_date":"2021-01-01 00:00:00"}]`
					w.WriteHeader(tc.mockStatusCode)
					w.Write([]byte(tc.mockResponse))
				} else if r.URL.Path == "/USD-EUR" {
					tc.mockResponse = `[{"code":"USD","codein":"EUR","name":"Dollar/Euro","high":"1.20","low":"1.10","varBid":"0.001","pctChange":"0.02","bid":"5.15","ask":"5.16","timestamp":"1609459200","create_date":"2021-01-01 00:00:00"}]`
					w.WriteHeader(tc.mockStatusCode)
					w.Write([]byte(tc.mockResponse))
				} else if tc.expectError {
					w.WriteHeader(tc.mockStatusCode)
					w.Write([]byte(tc.mockResponse))
				}
			}))
			defer mockServer.Close()
			currencyProvider := CurrencyProvider(CurrencyProviderImpl{
				OriginURL: mockServer.URL,
			})

			currencies, err := currencyProvider.GetCurrency(tc.code, tc.targets)

			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErrorResult, err.Error())
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tc.expectedResult, currencies)
			}
		})
	}
}
