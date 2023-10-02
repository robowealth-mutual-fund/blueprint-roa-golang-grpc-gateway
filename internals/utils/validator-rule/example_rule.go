package validatorrule

import (
	"github.com/go-playground/validator/v10"
)

// ExampleRule ...
type ExampleRule struct{ *CommonRule }

// GetRule ...
func (*ExampleRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		// To get field data fl.Field().String()

		return true
	}
}

// NewExampleRule ...
func NewExampleRule() *ExampleRule {
	return &ExampleRule{&CommonRule{
		Field:     "example",
		FieldName: "Example",
		Message:   "{0} invalid format!",
	}}
}
