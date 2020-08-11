package errors

import (
	"fmt"
	"net/http"
)

// UnauthorizedError used for when a user either cannot authenticate
// or cannot access a resource (cannot refresh idToken because of no refresh token)
// as an example
type UnauthorizedError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Error required for error interface
func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("UnauthorizedError")
}

// NewUnauthorized creates an error with status code 401
// and provides a message for the reason
func NewUnauthorized(reason string) *UnauthorizedError {
	return &UnauthorizedError{
		Type:    "Unauthorized",
		Message: reason,
		Status:  http.StatusUnauthorized,
	}
}
