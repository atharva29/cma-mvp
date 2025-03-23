package models

// Comparable represents a comparable property for CMA
// @Description A comparable property for CMA
type Comparable struct {
	// Property address
	// @Example 123 Main St
	Address string `json:"address"`

	// Sale price of the property
	// @Example 1100000
	SalePrice int `json:"sale_price"`

	// Square footage of the property
	// @Example 1300
	Sqft int `json:"sqft"`

	// Price per square foot
	// @Example 846
	PricePerSqft int `json:"price_per_sqft"`
}

// CMAResponse represents the comparative market analysis response
// @Description Comparative market analysis response
type CMAResponse struct {
	// Unique property identifier
	// @Example 12345
	PropertyID string `json:"property_id"`

	// List of comparable properties
	Comparables []Comparable `json:"comparables"`

	// Estimated property value based on comparables
	// @Example 1150000
	EstimatedValue int `json:"estimated_value"`
}

// CMARequest represents the request parameters for CMA
type CMARequest struct {
	PropertyID   string `json:"property_id"`
	Radius       int    `json:"radius"`
	PropertyType string `json:"property_type"`
}
