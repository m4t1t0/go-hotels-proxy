package home

import (
	"github.com/gofiber/fiber/v2"
)

func Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("{\"status\": \"ok\", \"version\": \"0.1\"}")
	}
}
