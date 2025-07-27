package Mapper

import (
	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server/handler/countries/model"
)

// CountryMapper defines the interface for mapping country data
type CountryMapper interface {
	// MapCountry maps a single raw country to a MappedCountry
	MapCountry(rawCountry interface{}, allCountries []interface{}) (model.MappedCountry, error)
	
	// MapCountries maps multiple raw countries to MappedCountry objects
	MapCountries(rawCountries []interface{}) ([]model.MappedCountry, error)
}