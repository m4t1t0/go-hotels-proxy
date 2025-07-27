package countries

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server/handler/countries/request"
)

// Handler returns a fiber.Handler that fetches countries from multiple regions concurrently
func Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a new countries service
		countriesService := request.NewCountriesService()
		
		// Handle the request using the countries service
		return countriesService.HandleCountriesRequest(c)
	}
}