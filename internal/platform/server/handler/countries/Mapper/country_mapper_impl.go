package Mapper

import (
	"errors"
	"fmt"
	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server/handler/countries/model"
)

// CountryMapperImpl implements the CountryMapper interface
type CountryMapperImpl struct{}

// NewCountryMapper creates a new CountryMapperImpl
func NewCountryMapper() CountryMapper {
	return &CountryMapperImpl{}
}

// MapCountry maps a single raw country to a MappedCountry
func (m *CountryMapperImpl) MapCountry(rawCountry interface{}, allCountries []interface{}) (model.MappedCountry, error) {
	// Type assertion for the raw country data
	countryMap, ok := rawCountry.(map[string]interface{})
	if !ok {
		return model.MappedCountry{}, errors.New("invalid country data format")
	}

	// Initialize the mapped country
	mappedCountry := model.MappedCountry{}

	// Map name.common to name field
	if nameData, ok := countryMap["name"].(map[string]interface{}); ok {
		if commonName, ok := nameData["common"].(string); ok {
			mappedCountry.Name = commonName
		} else {
			return model.MappedCountry{}, errors.New("name.common field is missing or not a string")
		}
	} else {
		return model.MappedCountry{}, errors.New("name field is missing or not a map")
	}

	// Map the key of the first node of currency to currency field
	if currencyData, ok := countryMap["currencies"].(map[string]interface{}); ok && len(currencyData) > 0 {
		// Get the first key from the currencies map
		for key := range currencyData {
			mappedCountry.Currency = key
			break
		}
	} else {
		// If currencies field is missing or empty, set currency to empty string
		mappedCountry.Currency = ""
	}

	// Map the first node of capital as a string to capital field
	if capitalData, ok := countryMap["capital"].([]interface{}); ok && len(capitalData) > 0 {
		if capital, ok := capitalData[0].(string); ok {
			mappedCountry.Capital = capital
		} else {
			// If first capital is not a string, set capital to empty string
			mappedCountry.Capital = ""
		}
	} else {
		// If capital field is missing or empty, set capital to empty string
		mappedCountry.Capital = ""
	}

	// Map the region node value to region field
	if region, ok := countryMap["region"].(string); ok {
		mappedCountry.Region = region
	} else {
		// If region field is missing or not a string, set region to empty string
		mappedCountry.Region = ""
	}

	// Map latlng node to coordinates field
	coordinates := model.Coordinates{}
	if latlngData, ok := countryMap["latlng"].([]interface{}); ok && len(latlngData) >= 2 {
		if lat, ok := latlngData[0].(float64); ok {
			coordinates.Latitude = lat
		}
		if lng, ok := latlngData[1].(float64); ok {
			coordinates.Longitude = lng
		}
	}
	mappedCountry.Coordinates = coordinates

	// Map borders field as an array of objects with name and capital fields
	borders := []model.Border{}
	if bordersData, ok := countryMap["borders"].([]interface{}); ok {
		for _, borderCode := range bordersData {
			if code, ok := borderCode.(string); ok {
				// Find the border country by cca3 code
				borderCountry := m.findCountryByCca3(code, allCountries)
				if borderCountry != nil {
					border := model.Border{}
					
					// Extract name from border country
					if nameData, ok := borderCountry["name"].(map[string]interface{}); ok {
						if commonName, ok := nameData["common"].(string); ok {
							border.Name = commonName
						}
					}
					
					// Extract capital from border country
					if capitalData, ok := borderCountry["capital"].([]interface{}); ok && len(capitalData) > 0 {
						if capital, ok := capitalData[0].(string); ok {
							border.Capital = capital
						}
					}
					
					borders = append(borders, border)
				}
			}
		}
	}
	mappedCountry.Borders = borders

	return mappedCountry, nil
}

// MapCountries maps multiple raw countries to MappedCountry objects
func (m *CountryMapperImpl) MapCountries(rawCountries []interface{}) ([]model.MappedCountry, error) {
	if rawCountries == nil {
		return nil, errors.New("raw countries data is nil")
	}

	mappedCountries := make([]model.MappedCountry, 0, len(rawCountries))
	
	for _, rawCountry := range rawCountries {
		mappedCountry, err := m.MapCountry(rawCountry, rawCountries)
		if err != nil {
			return nil, fmt.Errorf("error mapping country: %v", err)
		}
		mappedCountries = append(mappedCountries, mappedCountry)
	}
	
	return mappedCountries, nil
}

// findCountryByCca3 finds a country by its cca3 code
func (m *CountryMapperImpl) findCountryByCca3(cca3 string, countries []interface{}) map[string]interface{} {
	for _, country := range countries {
		countryMap, ok := country.(map[string]interface{})
		if !ok {
			continue
		}
		
		if code, ok := countryMap["cca3"].(string); ok && code == cca3 {
			return countryMap
		}
	}
	
	return nil
}