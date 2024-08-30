package validation

import (
	"encoding/json"
	"errors"

	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transl, _ = uni.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUser(
	validation_err error,
) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, err := range validation_err.(validator.ValidationErrors) {
			errorsCauses = append(errorsCauses, rest_err.Causes{
				Message: err.Translate(transl),
				Field:   err.Field(),
			})

		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)

	} else {
		return rest_err.NewBadRequestError("Invalid JSON")
	}
}
