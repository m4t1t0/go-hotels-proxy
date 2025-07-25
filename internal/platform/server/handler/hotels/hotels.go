package hotels

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hola %s", "Rodrigo"))
	}
}
