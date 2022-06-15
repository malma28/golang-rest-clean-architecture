package implement

import (
	"github.com/malma28/golang-rest-clean-architecture/entity"
	"github.com/malma28/golang-rest-clean-architecture/usecase/model"
	"github.com/malma28/golang-rest-clean-architecture/usecase/presenter"
)

type customerPresenterImplement struct {
}

func NewCustomerPresenter() presenter.CustomerPresenter {
	customerPresenter := new(customerPresenterImplement)

	return customerPresenter
}

func (presenter *customerPresenterImplement) Output(customerEntity entity.Customer) model.CustomerOutput {
	return model.CustomerOutput{
		Id:   customerEntity.Id,
		Name: customerEntity.Name,
	}
}
