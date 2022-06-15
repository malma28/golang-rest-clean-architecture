package validator

type Validator interface {
	Validate(data any) error
}
