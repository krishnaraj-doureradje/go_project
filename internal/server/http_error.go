package server

// ErrorResponse represents a standardized error response.
// It includes an application-specific error code and a human-readable message.
type ErrorResponse struct {
	// Code is a predefined error code string, e.g. "0001".
	Code string `json:"code" example:"0000" binding:"required"`

	// Message should clearly describe the error in human-readable form.
	Message string `json:"message" example:"The error in human-readable form" binding:"required"`
}
