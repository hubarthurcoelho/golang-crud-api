package validations

import "github.com/go-playground/validator/v10"

func ValidateBody(body any) error {
	bodyValidator := validator.New()

	if errs := bodyValidator.Struct(body); errs != nil {
		return errs
	}

	return nil
}
