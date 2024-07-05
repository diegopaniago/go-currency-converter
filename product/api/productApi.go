package api

import "github.com/gofiber/fiber/v3"

func AddProduct(c fiber.Ctx) error {
	return c.Send([]byte("I am alive!"))
}
