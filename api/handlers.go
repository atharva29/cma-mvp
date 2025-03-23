package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user/cma/models"
	"github.com/user/cma/modules"
)

// Handler struct contains all the dependencies for the API handlers
type Handler struct {
	marketAnalyzer *modules.MarketAnalyzer
	cmaAnalyzer    *modules.CMAAnalyzer
}

// NewHandler creates a new Handler instance
func NewHandler(marketAnalyzer *modules.MarketAnalyzer, cmaAnalyzer *modules.CMAAnalyzer) *Handler {
	return &Handler{
		marketAnalyzer: marketAnalyzer,
		cmaAnalyzer:    cmaAnalyzer,
	}
}

// GetMarketTrends handles the GET /market-trends endpoint
// @Summary Get real estate market trends
// @Description Fetches and analyzes real estate pricing trends for a specific location
// @ID get-market-trends
// @Produce json
// @Param location query string true "City, state, or ZIP code"
// @Param property_type query string false "Type of property (Single-family, condo, etc.)"
// @Param time_range query string false "Time range for analysis (e.g., Last 6 months, 1 year, etc.)" default(6 months)
// @Success 200 {object} models.MarketTrends
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /market-trends [get]
func (h *Handler) GetMarketTrends(c echo.Context) error {
	// Extract query parameters
	location := c.QueryParam("location")
	if location == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "location is required",
		})
	}

	propertyType := c.QueryParam("property_type")
	timeRange := c.QueryParam("time_range")
	if timeRange == "" {
		timeRange = "6 months" // Default to 6 months if not specified
	}

	// Create request model
	req := models.MarketTrendsRequest{
		Location:     location,
		PropertyType: propertyType,
		TimeRange:    timeRange,
	}

	// Get market trends
	trends, err := h.marketAnalyzer.GetMarketTrends(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch market trends: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, trends)
}

// GetCMA handles the GET /cma endpoint
// @Summary Get Comparative Market Analysis
// @Description Compares recent sales for a selected property to determine its market value
// @ID get-cma
// @Produce json
// @Param property_id query string true "Unique property identifier"
// @Param radius query integer false "Search radius in miles" default(5)
// @Param property_type query string false "Filter by property type"
// @Success 200 {object} models.CMAResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /cma [get]
func (h *Handler) GetCMA(c echo.Context) error {
	// Extract query parameters
	propertyID := c.QueryParam("property_id")
	if propertyID == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "property_id is required",
		})
	}

	radiusStr := c.QueryParam("radius")
	radius := 5 // Default radius is 5 miles
	if radiusStr != "" {
		var err error
		radius, err = strconv.Atoi(radiusStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "radius must be a valid integer",
			})
		}
	}

	propertyType := c.QueryParam("property_type")

	// Create request model
	req := models.CMARequest{
		PropertyID:   propertyID,
		Radius:       radius,
		PropertyType: propertyType,
	}

	// Get CMA
	cma, err := h.cmaAnalyzer.GetComparableProperties(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch CMA: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, cma)
}
