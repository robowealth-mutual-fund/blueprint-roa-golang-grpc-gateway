package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/constants"
	vr "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils/validator-rule"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

// CustomValidator ...
type CustomValidator struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

// Configure ...
func (cv *CustomValidator) Configure() {
	// Setup validator and initial language
	v := validator.New()

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	// Assign the Validate and Trans
	cv.Validator = v
	cv.Trans = trans

	// Start register a custom rules here ...
	vr.RegisterValidationRule(vr.NewExampleRule(), v, trans)
}

// Validate ...
func (cv *CustomValidator) Validate(structRule interface{}) error {
	if err := cv.Validator.Struct(structRule); err != nil {
		Error := NewError(constants.BAD_REQUEST, "Wrong Input")

		badRequest := &errdetails.BadRequest{}
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {

				jsonFieldName := e.Field()
				if field, ok := reflect.TypeOf(structRule).Elem().FieldByName(e.Field()); ok {
					if jsonTag, ok := field.Tag.Lookup("json"); ok {
						jsonFieldName = strings.Split(jsonTag, ",")[0]
					}
				}

				param := e.Param()
				if param != "" {
					param = "=" + param
				}

				badRequest.FieldViolations = append(badRequest.FieldViolations, &errdetails.BadRequest_FieldViolation{
					Field:       jsonFieldName,
					Description: fmt.Sprintf("ERROR_%s%s: %s", strings.ToUpper(e.Tag()), param, e.Translate(cv.Trans)),
				})
			}
		}

		Error.AddErrorDetails(badRequest)
		return Error.Err()
	}

	return nil
}

// NewCustomValidator ...
func NewCustomValidator() *CustomValidator {
	var cv = &CustomValidator{}
	cv.Configure()
	return cv
}
