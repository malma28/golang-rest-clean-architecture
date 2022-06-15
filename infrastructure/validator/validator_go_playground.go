package validator

import (
	"context"

	goplaygroundvalidator "github.com/go-playground/validator/v10"
	adaptervalidator "github.com/malma28/golang-rest-clean-architecture/adapter/validator"
)

type validatorGoPlayground struct {
	validate *goplaygroundvalidator.Validate
}

func newValidatorGoPlayground() adaptervalidator.Validator {
	validator := new(validatorGoPlayground)
	validator.validate = goplaygroundvalidator.New()

	return validator
}

func (validator *validatorGoPlayground) Validate(data any) error {
	return validator.validate.StructCtx(context.Background(), data)
}
