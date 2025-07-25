package request

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server/handler/coroutines/client"
	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server/handler/coroutines/model"
)

const (
	countriesAPIURL = "https://restcountries.com/v3.1/region/%s"
)

// CountriesService handles requests for countries data
type CountriesService struct {
	httpClient *client.HTTPClient
}

// NewCountriesService creates a new CountriesService
func NewCountriesService() *CountriesService {
	return &CountriesService{
		httpClient: client.NewHTTPClient(),
	}
}

// FetchCountriesByRegion fetches countries for a specific region
func (s *CountriesService) FetchCountriesByRegion(region string) ([]model.Country, error) {
	url := fmt.Sprintf(countriesAPIURL, region)
	
	body, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching countries for %s: %v", region, err)
	}
	
	var countries []model.Country
	if err := client.UnmarshalJSON(body, &countries); err != nil {
		return nil, fmt.Errorf("error parsing JSON for %s: %v", region, err)
	}
	
	return countries, nil
}

// FetchCountriesFromMultipleRegions fetches countries from multiple regions concurrently
func (s *CountriesService) FetchCountriesFromMultipleRegions(regions []string) ([]model.Country, error) {
	// Create a wait group to wait for all requests to complete
	var wg sync.WaitGroup
	wg.Add(len(regions))
	
	// Create a mutex to protect the combined results
	var mu sync.Mutex
	
	// Create a slice to store all countries
	var allCountries []model.Country
	
	// Create a channel to collect errors
	errCh := make(chan error, len(regions))
	
	// Fetch countries for each region concurrently
	for _, region := range regions {
		go func(reg string) {
			defer wg.Done()
			
			countries, err := s.FetchCountriesByRegion(reg)
			if err != nil {
				errCh <- err
				return
			}
			
			// Add the countries to the combined results
			mu.Lock()
			allCountries = append(allCountries, countries...)
			mu.Unlock()
		}(region)
	}
	
	// Wait for all requests to complete
	wg.Wait()
	close(errCh)
	
	// Check if there were any errors
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}
	
	return allCountries, nil
}

// HandleCountriesRequest handles the HTTP request for countries data
func (s *CountriesService) HandleCountriesRequest(c *fiber.Ctx) error {
	// Define the regions to fetch
	regions := []string{"europe", "africa"}
	
	// Fetch countries from multiple regions
	countries, err := s.FetchCountriesFromMultipleRegions(regions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	// Return the combined countries data
	return c.JSON(fiber.Map{
		"countries": countries,
	})
}