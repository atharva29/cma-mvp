package models

// MarketTrends represents real estate pricing trends for a specific location
// @Description Real estate pricing trends for a specific location
type MarketTrends struct {
	// Location (city, state, or ZIP code)
	// @Example San Francisco, CA
	Location string `json:"location"`

	// Median sale price in the area
	// @Example 1200000
	MedianPrice int `json:"median_price"`

	// Average price per square foot
	// @Example 900
	PricePerSqft int `json:"price_per_sqft"`

	// Number of sales in the given time period
	// @Example 120
	SalesVolume int `json:"sales_volume"`

	// Market trend direction (upward, downward, or stable)
	// @Example upward
	Trend string `json:"trend"`
}

// MarketTrendsRequest represents the request parameters for market trends
type MarketTrendsRequest struct {
	Location     string `json:"location"`
	PropertyType string `json:"property_type"`
	TimeRange    string `json:"time_range"`
}
