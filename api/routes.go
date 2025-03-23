package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @Summary Health check endpoint
// @Description Returns the current health status of the API
// @ID health-check
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"status": "healthy",
	})
}

// SetupRoutes configures all the routes for the API
func SetupRoutes(e *echo.Echo, h *Handler) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/market-trends", h.GetMarketTrends)
	e.GET("/cma", h.GetCMA)

	// Health check endpoint
	e.GET("/health", HealthCheck)

	// Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
