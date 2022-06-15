package usecase

import (
	"context"

	"github.com/malma28/golang-rest-clean-architecture/usecase/model"
)

type CustomerUsecase interface {
	AddCustomer(ctx context.Context, input model.AddCustomerInput) (model.CustomerOutput, error)
	GetAllCustomer(ctx context.Context) ([]model.CustomerOutput, error)
	GetCustomerById(ctx context.Context, id int64) (model.CustomerOutput, error)
	UpdateCustomerById(ctx context.Context, id int64, input model.UpdateCustomerInput) (model.CustomerOutput, error)
	DeleteCustomerById(ctx context.Context, id int64) error
}
