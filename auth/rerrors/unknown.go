package rerrors

import (
	"fmt"
)

// UnknownError holds a simple message as a final catch-all for errors
type UnknownError struct {
	Type   string `json:"type"`
	Status int    `json:"status"`
}

// Error required for error interface
func (e *UnknownError) Error() string {
	return fmt.Sprintf("UnknownError")
}

// NewUnknown creates an error with status code 409 and indicates the value that already exists
func NewUnknown(status int) *UnknownError {
	return &UnknownError{
		Type:   "Unknown",
		Status: status,
	}
}
