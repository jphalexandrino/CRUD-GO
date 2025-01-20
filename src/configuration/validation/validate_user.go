package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslation "github.com/go-playground/validator/v10/translations/en"
	"github.com/jphalexandrino/CRUD-GO/src/configuration/rest_err"
)

var (
	_      = validator.New()
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		eng := en.New()
		unt := ut.New(eng, eng)
		transl, _ = unt.GetTranslator("en")
		err := entranslation.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBedRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		var errorsCauses []rest_err.Causes

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBedRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBedRequestError("Error trying to convert fields")
	}

}
