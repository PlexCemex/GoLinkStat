package request

import "github.com/go-playground/validator/v10"

func ValidatorBody[T any](body T) error {
	validate := validator.New()
	err := validate.Struct(body)
	return err
}
