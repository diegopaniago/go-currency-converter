package service

import (
	"net/http"
	"testing"

	"github.com/diegopaniago/go-currency-converter/currency/model"
)

func TestGetCurrency(t *testing.T) {
	tests := []struct {
		name           string
		mockResponse   string
		mockStatusCode int
		code           string
		target         string
		expectError    bool
		expectedResult model.Currency
	}{
		{
			name:           "Currency Provider should return the currency",
			mockResponse:   `[{"code":"USD","codein":"BRL","name":"Dollar/Real","high":"5.20","low":"5.10","varBid":"0.001","pctChange":"0.02","bid":"5.15","ask":"5.16","timestamp":"1609459200","create_date":"2021-01-01 00:00:00"}]`,
			mockStatusCode: http.StatusOK,
			code:           "USD",
			target:         "BRL",
			expectError:    false,
			expectedResult: model.Currency{
				Code:  "USD",
				Name:  "Dollar",
				Price: 1.0000,
				Exchange: model.Exchange{
					Code:  "BRL",
					Name:  "Real",
					Price: 5.15,
				},
			},
		},
		{
			name:           "Currency Provider should return error if called with invalid code",
			mockResponse:   "",
			mockStatusCode: http.StatusInternalServerError,
			code:           "USD",
			target:         "BRL",
			expectError:    true,
		},
		{
			name:           "Currency Provider should return error if the external api returns with invalid JSON",
			mockResponse:   `invalid JSON`,
			mockStatusCode: http.StatusOK,
			code:           "USD",
			target:         "BRL",
			expectError:    true,
		},
	}

	for _, tc := range tests {
		// t.Run(tc.name, func(t *testing.T) {
		// 	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 		w.WriteHeader(tc.mockStatusCode)
		// 		w.Write([]byte(tc.mockResponse))
		// 	}))
		// 	defer mockServer.Close()

		// 	fmt.Println("mockServer.URL", mockServer.URL)

		// 	currencyProvider := CurrencyProvider(CurrencyProviderImpl{
		// 		OriginURL: mockServer.URL,
		// 	})
		// 	currency, err := currencyProvider.GetCurrency(tc.code, tc.target)

		// 	if tc.expectError {
		// 		assert.Error(t, err)
		// 	} else {
		// 		assert.NoError(t, err)
		// 		assert.Equal(t, tc.expectedResult, currency)
		// 	}
		// })
	}
}
