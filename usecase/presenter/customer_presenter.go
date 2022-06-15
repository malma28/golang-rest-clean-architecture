package presenter

import (
	"github.com/malma28/golang-rest-clean-architecture/entity"
	"github.com/malma28/golang-rest-clean-architecture/usecase/model"
)

type CustomerPresenter interface {
	Output(customerEntity entity.Customer) model.CustomerOutput
}
