# Comparative Market Analysis (CMA) API

A lightweight, API-only Comparative Market Analysis (CMA) and market trend tool built with Golang and Echo framework.

## Features

- Fetch and analyze real estate market trends
- Perform Comparative Market Analysis (CMA) for properties
- Clean API interface with JSON responses
- No database dependency (stateless)
- Interactive API documentation with Swagger UI

## API Documentation

The API documentation is available via Swagger UI at `/swagger` when the application is running.

You can access it at: http://localhost:8080/swagger

### API Endpoints

#### Get Market Trends
```
GET /market-trends

Query Parameters:
- location: City, state, or ZIP code
- property_type: Single-family, condo, etc.
- time_range: Last 6 months, 1 year, etc.
```

### Get Comparative Market Analysis (CMA)
```
GET /cma

Query Parameters:
- property_id: Unique property identifier
- radius: Search radius in miles
- property_type: Filter by property type
```

## Setup & Running

### Prerequisites
- Go 1.16+

### Installation
```bash
# Clone the repository
git clone https://github.com/user/cma.git
cd cma

# Install dependencies
go mod download

# Build the application
go build -o cma_api ./cmd/server

# Run the application
./cma_api
```

The API will be available at http://localhost:8080

## Environment Variables

- `PORT`: Port to run the server on (default: 8080)

## Development

```bash
# Run tests
go test ./...

# Run with hot reload (requires air)
air
```

## License

MIT # cma-mvp
