package model

// MappedCountry represents the transformed country data structure
type MappedCountry struct {
	Name        string      `json:"name"`
	Currency    string      `json:"currency"`
	Capital     string      `json:"capital"`
	Region      string      `json:"region"`
	Coordinates Coordinates `json:"coordinates"`
	Borders     []Border    `json:"borders"`
}

// Coordinates represents geographical coordinates
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Border represents a country that shares a border
type Border struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
}