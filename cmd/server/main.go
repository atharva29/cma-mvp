package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/user/cma/api"
	_ "github.com/user/cma/docs" // Import docs for swagger
	"github.com/user/cma/modules"
)

// @title Comparative Market Analysis (CMA) API
// @version 1.0
// @description A lightweight API for performing Comparative Market Analysis (CMA) and providing real estate market trends.
// @description This API fetches, analyzes, and returns structured real estate market data without using a database.

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

func main() {
	// Create echo instance
	e := echo.New()

	// Initialize dependencies
	dataFetcher := modules.NewDataFetcher()
	marketAnalyzer := modules.NewMarketAnalyzer(dataFetcher)
	cmaAnalyzer := modules.NewCMAAnalyzer(dataFetcher)

	// Create handler
	handler := api.NewHandler(marketAnalyzer, cmaAnalyzer)

	// Setup routes
	api.SetupRoutes(e, handler)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Starting server on port %s...", port)
	log.Printf("Swagger UI available at http://localhost:%s/swagger/index.html", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
