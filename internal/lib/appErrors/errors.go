package appErrors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidationErrors map[string]string

func (v ValidationErrors) Error() string {
	b := strings.Builder{}
	for field, value := range v {
		b.WriteString(field)
		b.WriteString(": ")
		b.WriteString(value)
		b.WriteString(";")
	}
	return b.String()
}

func NewValidationErrors() ValidationErrors {
	return make(ValidationErrors)
}

func NewValidationErrorsFromValidator(in validator.ValidationErrors) ValidationErrors {
	out := NewValidationErrors()
	for _, e := range in {
		field := strings.ToLower(e.Field())
		switch e.ActualTag() {
		case "required":
			out[field] = "field is a required"
		case "max":
			out[field] = fmt.Sprintf("max value %s", e.Param())
		case "min":
			out[field] = fmt.Sprintf("min value %s", e.Param())
		case "email":
			out[field] = "email is not valid"
		default:
			out[field] = "field is not valid"
		}
	}
	return out
}
