package api

import (
	"fmt"
	"strconv"

	"github.com/diegopaniago/go-currency-converter/product/service"
	"github.com/gofiber/fiber/v3"
)

var productService = service.IProductService(service.ProductService{})

func AddProduct(c fiber.Ctx) error {
	dto := new(service.ProductDto)
	c.Bind().Body(dto)
	var id, err = productService.AddProduct(*dto)
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return nil
	}
	c.Status(201).Send([]byte(fmt.Sprintf("%d", id)))
	return nil
}

func GetProduct(c fiber.Ctx) error {
	var id, err = strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(400).Send([]byte("Invalid id"))
		return nil
	}
	var dto, errNotFound = productService.GetProduct(id)
	if errNotFound != nil {
		c.Status(404).Send([]byte(errNotFound.Error()))
		return nil
	}
	c.Status(200).Send([]byte(fmt.Sprintf("%v", dto)))
	return nil
}
