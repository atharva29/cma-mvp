FROM golang:1.21-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cma_api ./cmd/server

# Use a small alpine image for the final container
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/cma_api .

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./cma_api"] 