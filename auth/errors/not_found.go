package errors

import (
	"fmt"
	"net/http"
)

// NotFoundError used for returning erros when resource entities cannot be found
// ... for instance in DB
type NotFoundError struct {
	Type            string          `json:"type"`
	InvalidArgument invalidArgument `json:"arg"`
	Status          int             `json:"status"`
}

// Error required for error interface
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("NotFound. The following resource:value could not be found: %v:%v\n", e.InvalidArgument.Name, e.InvalidArgument.Value)
}

// NewNotFound creates an error with status code 409 and indicates the value that already exists
func NewNotFound(name string, value string) *NotFoundError {
	return &NotFoundError{
		Type: "NotFound",
		InvalidArgument: invalidArgument{
			Name:  name,
			Value: value,
		},
		Status: http.StatusNotFound,
	}
}
