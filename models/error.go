package models

// ErrorResponse represents an error response from the API
// @Description Error response from the API
type ErrorResponse struct {
	// Error message
	// @Example location is required
	Error string `json:"error"`
}
