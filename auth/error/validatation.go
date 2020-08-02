package error

import "github.com/go-playground/validator/v10"

// ValidationError used to construct an error from validator.InvalidValidationError
// Which can can be used to conveniently send an error in a JSON response to the user
type ValidationError struct {
	Type        string
	InvalidArgs []string
}

// Error required for error interface
func (e *ValidationError) Error() string {
	return "ValidationError"
}

// NewFromValidationErrors constructs a ValidationError by formatting
// the properties of InvalidValidationError from the validator package
func NewFromValidationErrors(v validator.ValidationErrors) ValidationError {
	return ValidationError{
		Type:        "ValidationError",
		InvalidArgs: []string{},
	}
}
