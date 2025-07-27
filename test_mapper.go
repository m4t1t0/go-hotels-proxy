package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server/handler/countries/Mapper"
	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server/handler/countries/model"
)

func main() {
	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Define the URL for the REST Countries API (Europe region)
	url := "https://restcountries.com/v3.1/region/europe"

	// Make the HTTP request
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	// Decode the response body into a slice of interface{}
	var rawCountries []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawCountries); err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}

	// Create a new mapper
	mapper := Mapper.NewCountryMapper()

	// Map the raw countries data
	mappedCountries, err := mapper.MapCountries(rawCountries)
	if err != nil {
		log.Fatalf("Error mapping countries: %v", err)
	}

	// Find and print countries with borders
	fmt.Println("Looking for countries with borders:")
	
	// Countries that typically have borders
	targetCountries := []string{"Germany", "France", "Spain", "Italy"}
	
	// Print mapped countries with borders
	for _, country := range mappedCountries {
		for _, target := range targetCountries {
			if country.Name == target {
				fmt.Printf("\nFound %s with %d borders:\n", country.Name, len(country.Borders))
				printCountry(country)
				break
			}
		}
	}
	
	// Print the first 3 mapped countries
	fmt.Println("\nFirst 3 mapped countries:")
	for i, country := range mappedCountries {
		if i >= 3 {
			break
		}
		printCountry(country)
	}

	// Print the total number of mapped countries
	fmt.Printf("\nTotal mapped countries: %d\n", len(mappedCountries))
}

// printCountry prints the details of a mapped country
func printCountry(country model.MappedCountry) {
	fmt.Printf("\nCountry: %s\n", country.Name)
	fmt.Printf("  Currency: %s\n", country.Currency)
	fmt.Printf("  Capital: %s\n", country.Capital)
	fmt.Printf("  Region: %s\n", country.Region)
	fmt.Printf("  Coordinates: %.2f, %.2f\n", country.Coordinates.Latitude, country.Coordinates.Longitude)
	
	fmt.Printf("  Borders (%d):\n", len(country.Borders))
	for _, border := range country.Borders {
		fmt.Printf("    - %s (Capital: %s)\n", border.Name, border.Capital)
	}
}