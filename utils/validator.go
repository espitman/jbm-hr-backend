package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct using the validator package
// Returns nil if validation passes, otherwise returns a formatted validation error
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var errMsgs []string
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errMsgs = append(errMsgs, err.Field()+" is required")
		case "email":
			errMsgs = append(errMsgs, err.Field()+" must be a valid email address")
		case "oneof":
			errMsgs = append(errMsgs, err.Field()+" must be one of: "+err.Param())
		default:
			errMsgs = append(errMsgs, err.Field()+" failed on tag "+err.Tag())
		}
	}

	return errors.New(strings.Join(errMsgs, "; "))
}

// ValidateVar validates a single variable using the validator package
// Returns nil if validation passes, otherwise returns the validation error
func ValidateVar(field interface{}, tag string) error {
	return validate.Var(field, tag)
}
