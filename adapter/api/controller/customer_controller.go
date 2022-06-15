package controller

import (
	"io"

	"github.com/malma28/golang-rest-clean-architecture/adapter/api/response"
)

type CustomerController interface {
	CreateCustomer(body io.ReadCloser, contentType string) *response.ResponsePayload
	GetAllCustomer() *response.ResponsePayload
	GetCustomerById(id int64) *response.ResponsePayload
	UpdateCustomerById(id int64, body io.ReadCloser, contentType string) *response.ResponsePayload
	DeleteCustomerById(id int64) *response.ResponsePayload
}
