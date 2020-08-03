package error

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidationError used to construct an error from validator.InvalidValidationError
// Which can can be used to conveniently send an error in a JSON response to the user
type ValidationError struct {
	Type        string            `json:"type"`
	InvalidArgs []invalidArgument `json:"invalidArgs"`
}

type invalidArgument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Error required for error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("ValidationError. The following arguments and values are invalid: %+v\n", e.InvalidArgs)
}

// NewFromValidationErrors constructs a ValidationError by formatting
// the properties of InvalidValidationError from the validator package
func NewFromValidationErrors(vs validator.ValidationErrors) ValidationError {
	var invalidArgs []invalidArgument

	for _, err := range vs {
		invalidArgs = append(invalidArgs, invalidArgument{
			Name:  err.Field(),
			Value: err.Value().(string),
		})
	}

	return ValidationError{
		Type:        "ValidationError",
		InvalidArgs: invalidArgs,
	}
}
