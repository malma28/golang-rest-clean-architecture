package interactor

import (
	"context"

	"github.com/malma28/golang-rest-clean-architecture/adapter/repository"
	"github.com/malma28/golang-rest-clean-architecture/entity"
	"github.com/malma28/golang-rest-clean-architecture/usecase"
	"github.com/malma28/golang-rest-clean-architecture/usecase/model"
	"github.com/malma28/golang-rest-clean-architecture/usecase/presenter"
)

type customerInteractor struct {
	customerRepository repository.CustomerRepository
	customerPresenter  presenter.CustomerPresenter
}

func NewCustomerInteractor(customerRepository repository.CustomerRepository, customerPresenter presenter.CustomerPresenter) usecase.CustomerUsecase {
	customerInteractor := new(customerInteractor)

	customerInteractor.customerRepository = customerRepository
	customerInteractor.customerPresenter = customerPresenter

	return customerInteractor
}

func (interactor *customerInteractor) AddCustomer(ctx context.Context, input model.AddCustomerInput) (model.CustomerOutput, error) {
	output := model.CustomerOutput{}

	customerEntity, err := interactor.customerRepository.Save(ctx, entity.Customer{Name: input.Name})
	if err != nil {
		return output, err
	}

	output = interactor.customerPresenter.Output(customerEntity)

	return output, nil
}

func (interactor *customerInteractor) GetAllCustomer(ctx context.Context) ([]model.CustomerOutput, error) {
	output := []model.CustomerOutput{}

	customersEntity, err := interactor.customerRepository.FindAll(ctx)
	if err != nil {
		return output, err
	}

	for _, customerEntity := range customersEntity {
		output = append(output, interactor.customerPresenter.Output(customerEntity))
	}

	return output, nil
}

func (interactor *customerInteractor) GetCustomerById(ctx context.Context, id int64) (model.CustomerOutput, error) {
	output := model.CustomerOutput{}

	customerEntity, err := interactor.customerRepository.FindById(ctx, id)
	if err != nil {
		return output, err
	}

	output = interactor.customerPresenter.Output(customerEntity)

	return output, nil
}

func (interactor *customerInteractor) UpdateCustomerById(ctx context.Context, id int64, input model.UpdateCustomerInput) (model.CustomerOutput, error) {
	output := model.CustomerOutput{}

	customerEntity, err := interactor.customerRepository.UpdateById(
		ctx,
		id,
		entity.Customer{
			Name: input.Name,
		},
	)
	if err != nil {
		return output, err
	}

	output = interactor.customerPresenter.Output(customerEntity)

	return output, nil
}

func (interactor *customerInteractor) DeleteCustomerById(ctx context.Context, id int64) error {
	return interactor.customerRepository.DeleteById(
		ctx,
		id,
	)
}
