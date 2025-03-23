package modules

import (
	"github.com/user/cma/models"
)

// CMAAnalyzer handles the Comparative Market Analysis
type CMAAnalyzer struct {
	dataFetcher *DataFetcher
}

// NewCMAAnalyzer creates a new CMAAnalyzer instance
func NewCMAAnalyzer(df *DataFetcher) *CMAAnalyzer {
	return &CMAAnalyzer{
		dataFetcher: df,
	}
}

// GetComparableProperties fetches comparable properties for a given property ID
func (ca *CMAAnalyzer) GetComparableProperties(req models.CMARequest) (*models.CMAResponse, error) {
	// In a real implementation, we would:
	// 1. Fetch details of the target property
	// 2. Search for recently sold properties with similar characteristics in the given radius
	// 3. Calculate the estimated value based on comparables

	// For now, we're returning mock data
	// This can be extended to use real data sources

	// Mocked comparables for demonstration
	comparables := []models.Comparable{
		{
			Address:      "123 Main St",
			SalePrice:    1100000,
			Sqft:         1300,
			PricePerSqft: 846,
		},
		{
			Address:      "456 Elm St",
			SalePrice:    1150000,
			Sqft:         1400,
			PricePerSqft: 821,
		},
		{
			Address:      "789 Oak St",
			SalePrice:    1200000,
			Sqft:         1380,
			PricePerSqft: 870,
		},
	}

	// Calculate estimated value (average of comparable prices)
	var totalPrice int
	for _, comp := range comparables {
		totalPrice += comp.SalePrice
	}

	estimatedValue := 0
	if len(comparables) > 0 {
		estimatedValue = totalPrice / len(comparables)
	}

	return &models.CMAResponse{
		PropertyID:     req.PropertyID,
		Comparables:    comparables,
		EstimatedValue: estimatedValue,
	}, nil
}

// CalculatePricePerSqft calculates the price per square foot for a property
func (ca *CMAAnalyzer) CalculatePricePerSqft(price, sqft int) int {
	if sqft <= 0 {
		return 0
	}
	return price / sqft
}
