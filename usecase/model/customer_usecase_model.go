package model

type AddCustomerInput struct {
	Name string `json:"name" validate:"required,gte=1"`
}

type UpdateCustomerInput struct {
	Name string `json:"name" validate:"gte=1"`
}

type CustomerOutput struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
