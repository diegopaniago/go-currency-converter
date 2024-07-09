package service

import (
	"fmt"

	"github.com/diegopaniago/go-currency-converter/product/domain"
)

type IProductService interface {
	AddProduct(dto ProductDto) (int, error)
	GetProduct(id int) (ProductDto, error)
}

type ProductService struct{}

func (ProductService) AddProduct(dto ProductDto) (int, error) {
	var _, err = domain.NewProduct(dto.Name, dto.Price, dto.Currencies)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (ProductService) GetProduct(id int) (ProductDto, error) {
	return ProductDto{}, fmt.Errorf("not implemented")
}
