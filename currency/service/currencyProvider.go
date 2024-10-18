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
	GetCurrency(code string, targets []string) ([]model.Currency, error)
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

type CurrencyTask struct {
	Currency model.Currency
	Err      error
}

func (c CurrencyProviderImpl) GetCurrency(code string, targets []string) ([]model.Currency, error) {
	currencyTasks := make(chan CurrencyTask)
	var currencies []model.Currency

	for _, target := range targets {
		go c.fetchCurrency(code, target, currencyTasks)
	}

	for i := 0; i < len(targets); i++ {
		currencyTask := <-currencyTasks
		if currencyTask.Err != nil {
			return nil, currencyTask.Err
		} else {
			currencies = append(currencies, currencyTask.Currency)
		}
	}

	return currencies, nil
}

func (c CurrencyProviderImpl) fetchCurrency(code string, target string, currencyTasks chan<- CurrencyTask) {
	url := fmt.Sprintf(c.OriginURL+"/%s-%s", code, target)
	res, err := http.Get(url)

	if err != nil {
		currencyTasks <- CurrencyTask{Err: err}
		return
	}
	if res.StatusCode != http.StatusOK {
		currencyTasks <- CurrencyTask{Err: fmt.Errorf("error fetching currency %s to %s: %v", code, target, res.Status)}
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		currencyTasks <- CurrencyTask{Err: err}
		return
	}
	var cotations []CotationResponse
	if err := json.Unmarshal(body, &cotations); err != nil {
		currencyTasks <- CurrencyTask{Err: fmt.Errorf("error to unmarshal json: %v", err)}
		return
	}
	cotationPrice, err := strconv.ParseFloat(cotations[0].Bid, 64)
	if err != nil {
		currencyTasks <- CurrencyTask{Err: err}
		return
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
	currencyTasks <- CurrencyTask{Currency: currency, Err: nil}
}
