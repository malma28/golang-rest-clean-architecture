package validator

import "github.com/malma28/golang-rest-clean-architecture/adapter/validator"

type ValidatorType int

const (
	ValidatorGoPlayground ValidatorType = iota
)

func NewValidator(validatorType ValidatorType) validator.Validator {
	switch validatorType {
	case ValidatorGoPlayground:
		return newValidatorGoPlayground()
	}
	return nil
}
