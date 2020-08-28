package rerrors

import (
	"fmt"
	"net/http"
)

// AlreadyExistsError holds the field of a resource that failed to be created
// because the field already exists
type AlreadyExistsError struct {
	Type            string          `json:"type"`
	InvalidArgument invalidArgument `json:"arg"`
	Status          int             `json:"status"`
}

// Error required for error interface
func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("AlreadyExistsError. The following resource already exists: %v\n", e.InvalidArgument.Name)
}

// NewAlreadyExists creates an error with status code 409 and indicates the value that already exists
func NewAlreadyExists(name string, value string) *AlreadyExistsError {
	return &AlreadyExistsError{
		Type: "AlreadyExists",
		InvalidArgument: invalidArgument{
			Name:  name,
			Value: value,
		},
		Status: http.StatusConflict,
	}
}
