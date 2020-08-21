package errors

import (
	"fmt"
	"net/http"
)

// ExceedsMaxSizeError holds the field of a resource that failed to be created
// because the field already exists
type ExceedsMaxSizeError struct {
	Type                 string `json:"type"`
	Status               int    `json:"status"`
	ContentLength        int64  `json:"contentLength"`
	MaxAllowableBodySize int64  `json:"maxAllowableBodySize"`
}

// Error required for error interface
func (e *ExceedsMaxSizeError) Error() string {
	return fmt.Sprintf("ExceedsMaxSizeError. Request body exceeds maximum allowable size of: %v MB\n", e.MaxAllowableBodySize)
}

// NewExceedsMaxSize creates an error with status code 409 and indicates the value that already exists
func NewExceedsMaxSize(maxBodySize int64, contentLength int64) *ExceedsMaxSizeError {
	return &ExceedsMaxSizeError{
		Type:                 "ExceedsMaxSize",
		Status:               http.StatusRequestEntityTooLarge,
		MaxAllowableBodySize: maxBodySize,
		ContentLength:        contentLength,
	}
}
