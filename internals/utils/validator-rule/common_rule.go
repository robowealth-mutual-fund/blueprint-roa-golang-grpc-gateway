package validatorrule

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Rule ...
type Rule interface {
	RegisterMessageTranslate(ut.Translator) error
	RegisterFieldNameTranslate(ut.Translator, validator.FieldError) string
	GetField() string
	GetFieldName() string
	GetMessage() string
	GetRule() func(fl validator.FieldLevel) bool
}

// CommonRule ...
type CommonRule struct {
	Field     string
	FieldName string
	Message   string
}

// GetField ...
func (rule *CommonRule) GetField() string {
	return rule.Field
}

// GetFieldName ...
func (rule *CommonRule) GetFieldName() string {
	return rule.FieldName
}

// GetMessage ...
func (rule *CommonRule) GetMessage() string {
	return rule.Message
}

// RegisterMessageTranslate ...
func (rule *CommonRule) RegisterMessageTranslate(ut ut.Translator) error {
	return ut.Add(rule.Field, rule.Message, true)
}

// RegisterFieldNameTranslate ...
func (rule *CommonRule) RegisterFieldNameTranslate(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T(rule.Field, rule.FieldName)
	return t
}

// RegisterValidationRule ...
func RegisterValidationRule(rule Rule, validate *validator.Validate, trans ut.Translator) {
	validate.RegisterValidation(rule.GetField(), rule.GetRule())
	validate.RegisterTranslation(rule.GetField(), trans, rule.RegisterMessageTranslate, rule.RegisterFieldNameTranslate)
}
