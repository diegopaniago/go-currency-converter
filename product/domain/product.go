package domain

import (
	"fmt"
	"strings"
)

type Product struct {
	id         int
	name       string
	price      float32
	currencies []string
}

func NewProduct(name string, price float32, currencies []string) (Product, error) {
	var validationErros []string = []string{}
	if name == "" {
		validationErros = append(validationErros, "Name is required")
	}
	if price <= 0 {
		validationErros = append(validationErros, "Price must be greater than zero")
	}
	if len(currencies) == 0 {
		validationErros = append(validationErros, "At least one currency is required")
	}
	if len(validationErros) > 0 {
		return Product{}, fmt.Errorf(strings.Join(validationErros, ", "))
	}
	return Product{name: name, price: price, currencies: currencies}, nil
}
