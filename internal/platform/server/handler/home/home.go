package home

import (
	"github.com/gofiber/fiber/v2"
)

func Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Ola ke ase?, eta en el handler o ke ase?")
	}
}
