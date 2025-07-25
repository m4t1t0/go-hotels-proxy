package model

// Country represents a country data structure
// Using interface{} to match the original implementation
// In a real-world scenario, we would define a proper struct with specific fields
type Country interface{}

// CountriesResponse represents the response structure for the countries API
type CountriesResponse struct {
	Countries []Country `json:"countries"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}