package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/diegopaniago/go-currency-converter/currency/model"
)

type CurrencyProvider interface {
	GetCurrency(code string, target string) (model.Currency, error)
}

type CurrencyProviderImpl struct{}

type CotationResponse struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func (CurrencyProviderImpl) GetCurrency(code string, target string) (model.Currency, error) {
	res, err := http.Get(fmt.Sprintf("https://economia.awesomeapi.com.br/%s-%s", code, target))
	if err != nil {
		return model.Currency{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.Currency{}, err
	}
	var cotations []CotationResponse
	if err := json.Unmarshal(body, &cotations); err != nil {
		return model.Currency{}, err
	}
	cotationPrice, err := strconv.ParseFloat(cotations[0].Bid, 64)
	if err != nil {
		return model.Currency{}, err
	}
	currency := model.Currency{
		Code:  cotations[0].Code,
		Name:  strings.Split(cotations[0].Name, "/")[0],
		Price: 1.0000,
		Exchange: model.Exchange{
			Code:  cotations[0].Codein,
			Name:  strings.Split(cotations[0].Name, "/")[1],
			Price: cotationPrice,
		},
	}
	return currency, nil
}
