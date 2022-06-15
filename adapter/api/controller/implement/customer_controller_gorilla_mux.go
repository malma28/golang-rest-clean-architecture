package implement

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/malma28/golang-rest-clean-architecture/adapter/api/controller"
	"github.com/malma28/golang-rest-clean-architecture/adapter/api/response"
	"github.com/malma28/golang-rest-clean-architecture/adapter/validator"
	"github.com/malma28/golang-rest-clean-architecture/entity/exception"
	"github.com/malma28/golang-rest-clean-architecture/usecase"
	"github.com/malma28/golang-rest-clean-architecture/usecase/model"
)

type customerControllerGorillaMux struct {
	validator       validator.Validator
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerController(validator validator.Validator, customerUsecase usecase.CustomerUsecase) controller.CustomerController {
	customerController := new(customerControllerGorillaMux)
	customerController.validator = validator
	customerController.customerUsecase = customerUsecase

	return customerController
}

func (controller *customerControllerGorillaMux) CreateCustomer(body io.ReadCloser, contentType string) *response.ResponsePayload {
	res := new(response.ResponsePayload)

	if contentType != "application/json" {
		return res.SetData(nil).SetMessage("Unsupported Media Type").SetStatus(http.StatusUnsupportedMediaType).SetSuccess(false)
	}

	input := model.AddCustomerInput{}

	defer body.Close()

	if err := json.NewDecoder(body).Decode(&input); err != nil {
		return res.SetData(nil).SetMessage("Bad Request").
			SetStatus(http.StatusBadRequest).SetSuccess(false)
	}

	if err := controller.validator.Validate(input); err != nil {
		return res.SetData(nil).SetMessage("Bad Request").
			SetStatus(http.StatusBadRequest).SetSuccess(false)
	}

	output, err := controller.customerUsecase.AddCustomer(context.Background(), input)
	if err != nil {
		return controller.handleError(err)
	}

	return res.SetData(output).SetMessage("Created").SetStatus(http.StatusCreated).SetSuccess(true)
}

func (controller *customerControllerGorillaMux) GetAllCustomer() *response.ResponsePayload {
	res := new(response.ResponsePayload)

	outputs, err := controller.customerUsecase.GetAllCustomer(context.Background())
	if err != nil {
		return controller.handleError(err)
	}

	if len(outputs) < 1 {
		return res.SetMessage("Not Found").SetStatus(http.StatusNotFound).SetSuccess(true)
	}

	return res.SetData(outputs).SetMessage("OK").SetStatus(http.StatusOK).SetSuccess(true)
}

func (controller *customerControllerGorillaMux) GetCustomerById(id int64) *response.ResponsePayload {
	res := new(response.ResponsePayload)

	output, err := controller.customerUsecase.GetCustomerById(context.Background(), id)
	if err != nil {
		return controller.handleError(err)
	}

	return res.SetData(output).SetMessage("OK").SetStatus(http.StatusOK).SetSuccess(true)
}

func (controller *customerControllerGorillaMux) UpdateCustomerById(id int64, body io.ReadCloser, contentType string) *response.ResponsePayload {
	res := new(response.ResponsePayload)

	if contentType != "application/json" {
		return res.SetData(nil).SetMessage("Unsupported Media Type").SetStatus(http.StatusUnsupportedMediaType).SetSuccess(false)
	}

	input := model.UpdateCustomerInput{}

	defer body.Close()

	if err := json.NewDecoder(body).Decode(&input); err != nil {
		return res.SetData(nil).SetMessage("Bad Request").
			SetStatus(http.StatusBadRequest).SetSuccess(false)
	}

	if err := controller.validator.Validate(input); err != nil {
		return res.SetData(nil).SetMessage("Bad Request").
			SetStatus(http.StatusBadRequest).SetSuccess(false)
	}

	output, err := controller.customerUsecase.UpdateCustomerById(context.Background(), id, input)
	if err != nil {
		return controller.handleError(err)
	}

	return res.SetData(output).SetMessage("OK").SetStatus(http.StatusOK).SetSuccess(true)
}

func (controller *customerControllerGorillaMux) DeleteCustomerById(id int64) *response.ResponsePayload {
	res := new(response.ResponsePayload)

	err := controller.customerUsecase.DeleteCustomerById(context.Background(), id)
	if err != nil {
		return controller.handleError(err)
	}

	return res.SetData(nil).SetMessage("OK").SetStatus(http.StatusOK).SetSuccess(true)
}

func (customerController *customerControllerGorillaMux) handleError(err error) *response.ResponsePayload {
	switch err {
	case exception.ErrCustomerNotFound:
		return response.FromError(err).SetStatus(http.StatusNotFound)
	}
	return response.FromError(err).SetMessage("Internal Server Error")
}
