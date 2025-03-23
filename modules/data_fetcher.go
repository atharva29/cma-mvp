package modules

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// DataFetcher handles fetching data from external API sources
type DataFetcher struct {
	client *http.Client
}

// NewDataFetcher creates a new DataFetcher instance
func NewDataFetcher() *DataFetcher {
	return &DataFetcher{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// FetchJSON fetches JSON data from a URL and unmarshals it into the provided interface
func (df *DataFetcher) FetchJSON(url string, target interface{}) error {
	resp, err := df.client.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return nil
}

// FetchMockData returns mock data for testing purposes
func (df *DataFetcher) FetchMockData() ([]byte, error) {
	// This function will return mock data for development and testing
	// In a real implementation, we would connect to real estate APIs
	return []byte(`{"status": "success", "data": {}}`), nil
}
