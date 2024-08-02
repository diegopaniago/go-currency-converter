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
	GetCurrency(code string, targets []string) (model.Currency, error)
}

type CurrencyProviderImpl struct {
	OriginURL string
}

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

func (c CurrencyProviderImpl) GetCurrency(code string, targets []string) (model.Currency, error) {
	//TODO: Refact to use goroutines to get all targets at the same time
	target := strings.Join(targets, ",")
	url := fmt.Sprintf(c.OriginURL+"/%s-%s", code, target)
	res, err := http.Get(url)
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
