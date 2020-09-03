package rerrors

// package rerrors shares errors accross
// all application layers.
// In a more production grade application, each layer should probably
// have its own error type which can then be handled with
// the layer down the request chain

import (
	"errors"
	"fmt"
	"net/http"
)

// Type holds a type string and integer code for the error
type Type string

// Set of valid errorTypes
const (
	Authorization   Type = "AUTHORIZATION"   // Authentication Failures -
	BadRequest      Type = "BADREQUEST"      // Validation errors / BadInput
	Conflict        Type = "CONFLICT"        // Already exists (eg, create account with existent email) - 409
	Internal        Type = "INTERNAL"        // Server (500) and fallback errors
	NotFound        Type = "NOTFOUND"        // For not finding resource
	PayloadTooLarge Type = "PAYLOADTOOLARGE" // for uploading tons of JSON, or an image over the limit - 413
)

// Error holds custom error types/codes using graphql-go error extensions interface
type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

// Error satisfies standard error interface
func (e *Error) Error() string {
	return e.Message
}

// Status is a mapping errors to status codes
// Of course, this is somewhat redundant since
// our errors already map http status codes
func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusInternalServerError
	}
}

// Status checks the runtime type
// of the error and returns an http
// status code if the error is rerrors.Error
func Status(err error) int {
	var rerr *Error
	if errors.As(err, &rerr) {
		return rerr.Status()
	}
	return http.StatusInternalServerError
}

/*
* Error "Factories"
 */

// NewAuthorization to create a 401
func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

// NewBadRequest to create 400 errors (validation, for example)
func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

// NewConflict to create an error for 409
func NewConflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

// NewInternal for 500 errors and unknown errors
func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error."),
	}
}

// NewNotFound to create an error for 404
func NewNotFound(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("resource: %v with value: %v not found", name, value),
	}
}

// NewPayloadTooLarge to create an error for 413
func NewPayloadTooLarge(maxBodySize int64, contentLength int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
	}
}
