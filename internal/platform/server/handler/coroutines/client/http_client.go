package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HTTPClient is a wrapper around the standard http.Client
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTPClient
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

// Get performs a GET request to the specified URL
func (c *HTTPClient) Get(url string) ([]byte, error) {
	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Send the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-OK status: %d, body: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// UnmarshalJSON unmarshals JSON data into the provided interface
func UnmarshalJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}