package modules

import (
	"testing"

	"github.com/user/cma/models"
)

func TestAnalyzeTrend(t *testing.T) {
	// Create test cases
	testCases := []struct {
		name           string
		historicalData []float64
		expected       string
	}{
		{
			name:           "Upward Trend",
			historicalData: []float64{100, 105, 110, 115, 120},
			expected:       "upward",
		},
		{
			name:           "Downward Trend",
			historicalData: []float64{120, 115, 110, 105, 100},
			expected:       "downward",
		},
		{
			name:           "Stable Trend",
			historicalData: []float64{100, 102, 101, 103, 102},
			expected:       "stable",
		},
		{
			name:           "Insufficient Data",
			historicalData: []float64{100},
			expected:       "stable",
		},
		{
			name:           "Empty Data",
			historicalData: []float64{},
			expected:       "stable",
		},
	}

	// Create dependencies
	df := NewDataFetcher()
	analyzer := NewMarketAnalyzer(df)

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := analyzer.AnalyzeTrend(tc.historicalData)
			if result != tc.expected {
				t.Errorf("Expected %s but got %s", tc.expected, result)
			}
		})
	}
}

func TestGetMarketTrends(t *testing.T) {
	// Create dependencies
	df := NewDataFetcher()
	analyzer := NewMarketAnalyzer(df)

	// Test request
	req := models.MarketTrendsRequest{
		Location:     "San Francisco, CA",
		PropertyType: "Single-family",
		TimeRange:    "6 months",
	}

	// Call the function
	result, err := analyzer.GetMarketTrends(req)

	// Check for errors
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	// Check that result is not nil
	if result == nil {
		t.Error("Expected non-nil result but got nil")
		return
	}

	// Check that location matches request
	if result.Location != req.Location {
		t.Errorf("Expected location %s but got %s", req.Location, result.Location)
	}

	// Check that other fields have reasonable values
	if result.MedianPrice <= 0 {
		t.Errorf("Expected positive median price but got %d", result.MedianPrice)
	}

	if result.PricePerSqft <= 0 {
		t.Errorf("Expected positive price per sqft but got %d", result.PricePerSqft)
	}

	if result.SalesVolume <= 0 {
		t.Errorf("Expected positive sales volume but got %d", result.SalesVolume)
	}

	if result.Trend == "" {
		t.Error("Expected non-empty trend but got empty string")
	}
}
