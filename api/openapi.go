package api

// OpenAPISpec is the OpenAPI specification for the CMA API
const OpenAPISpec = `
openapi: 3.0.0
info:
  title: Comparative Market Analysis (CMA) API
  description: |
    A lightweight API for performing Comparative Market Analysis (CMA) and providing real estate market trends.
    This API fetches, analyzes, and returns structured real estate market data without using a database.
  version: 1.0.0
  contact:
    name: API Support
    email: support@example.com

servers:
  - url: http://localhost:8080
    description: Local development server

paths:
  /market-trends:
    get:
      summary: Get real estate market trends
      description: Fetches and analyzes real estate pricing trends for a specific location
      operationId: getMarketTrends
      parameters:
        - name: location
          in: query
          required: true
          description: City, state, or ZIP code
          schema:
            type: string
          example: San Francisco, CA
        - name: property_type
          in: query
          required: false
          description: Type of property (Single-family, condo, etc.)
          schema:
            type: string
          example: Single-family
        - name: time_range
          in: query
          required: false
          description: Time range for analysis (e.g., Last 6 months, 1 year, etc.)
          schema:
            type: string
            default: 6 months
          example: 1 year
      responses:
        200:
          description: Market trends data retrieved successfully
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/MarketTrends'
              examples:
                sanFrancisco:
                  summary: Market trend data for San Francisco
                  value:
                    location: San Francisco, CA
                    median_price: 1200000
                    price_per_sqft: 900
                    sales_volume: 120
                    trend: upward
                newYork:
                  summary: Market trend data for New York City
                  value:
                    location: New York, NY
                    median_price: 950000
                    price_per_sqft: 800
                    sales_volume: 200
                    trend: stable
        400:
          description: Bad request - missing required parameters
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
              example:
                error: location is required
        500:
          description: Internal server error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
              example:
                error: failed to fetch market trends

  /cma:
    get:
      summary: Get Comparative Market Analysis
      description: |
        Compares recent sales for a selected property to determine its market value.
        Returns comparable properties and an estimated property value.
      operationId: getCMA
      parameters:
        - name: property_id
          in: query
          required: true
          description: Unique property identifier
          schema:
            type: string
          example: 12345
        - name: radius
          in: query
          required: false
          description: Search radius in miles
          schema:
            type: integer
            default: 5
          example: 3
        - name: property_type
          in: query
          required: false
          description: Filter by property type
          schema:
            type: string
          example: Single-family
      responses:
        200:
          description: CMA data retrieved successfully
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/CMAResponse'
              examples:
                singleFamilyHome:
                  summary: CMA for a single-family home
                  value:
                    property_id: "12345"
                    comparables:
                      - address: "123 Main St"
                        sale_price: 1100000
                        sqft: 1300
                        price_per_sqft: 846
                      - address: "456 Elm St"
                        sale_price: 1150000
                        sqft: 1400
                        price_per_sqft: 821
                      - address: "789 Oak St"
                        sale_price: 1200000
                        sqft: 1380
                        price_per_sqft: 870
                    estimated_value: 1150000
                condo:
                  summary: CMA for a condominium
                  value:
                    property_id: "67890"
                    comparables:
                      - address: "101 Tower Ave, #405"
                        sale_price: 750000
                        sqft: 900
                        price_per_sqft: 833
                      - address: "101 Tower Ave, #512"
                        sale_price: 780000
                        sqft: 925
                        price_per_sqft: 843
                      - address: "202 High Rise Blvd, #301"
                        sale_price: 760000
                        sqft: 910
                        price_per_sqft: 835
                    estimated_value: 763333
        400:
          description: Bad request - missing required parameters
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
              example:
                error: property_id is required
        500:
          description: Internal server error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
              example:
                error: failed to fetch CMA

  /health:
    get:
      summary: Health check endpoint
      description: Returns the current health status of the API
      operationId: healthCheck
      responses:
        200:
          description: Service is healthy
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
              example:
                status: healthy

components:
  schemas:
    MarketTrends:
      type: object
      required:
        - location
        - median_price
        - price_per_sqft
        - sales_volume
        - trend
      properties:
        location:
          type: string
          description: The location (city, state, or ZIP code)
          example: San Francisco, CA
        median_price:
          type: integer
          description: The median sale price in the area
          example: 1200000
        price_per_sqft:
          type: integer
          description: The average price per square foot
          example: 900
        sales_volume:
          type: integer
          description: The number of sales in the given time period
          example: 120
        trend:
          type: string
          description: Market trend direction (upward, downward, or stable)
          enum:
            - upward
            - downward
            - stable
          example: upward

    Comparable:
      type: object
      required:
        - address
        - sale_price
        - sqft
        - price_per_sqft
      properties:
        address:
          type: string
          description: Property address
          example: 123 Main St
        sale_price:
          type: integer
          description: Sale price of the property
          example: 1100000
        sqft:
          type: integer
          description: Square footage of the property
          example: 1300
        price_per_sqft:
          type: integer
          description: Price per square foot
          example: 846

    CMAResponse:
      type: object
      required:
        - property_id
        - comparables
        - estimated_value
      properties:
        property_id:
          type: string
          description: Unique property identifier
          example: "12345"
        comparables:
          type: array
          description: List of comparable properties
          items:
            $ref: '#/components/schemas/Comparable'
        estimated_value:
          type: integer
          description: Estimated property value based on comparables
          example: 1150000

    Error:
      type: object
      required:
        - error
      properties:
        error:
          type: string
          description: Error message
          example: Internal server error
`
