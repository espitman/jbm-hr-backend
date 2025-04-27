package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct using the validator package
// Returns nil if validation passes, otherwise returns the validation error
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// ValidateVar validates a single variable using the validator package
// Returns nil if validation passes, otherwise returns the validation error
func ValidateVar(field interface{}, tag string) error {
	return validate.Var(field, tag)
}
