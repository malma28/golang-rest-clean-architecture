package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	controllerimplement "github.com/malma28/golang-rest-clean-architecture/adapter/api/controller/implement"
	"github.com/malma28/golang-rest-clean-architecture/adapter/api/response"
	"github.com/malma28/golang-rest-clean-architecture/adapter/database"
	"github.com/malma28/golang-rest-clean-architecture/adapter/repository/mysql"
	"github.com/malma28/golang-rest-clean-architecture/adapter/validator"
	"github.com/malma28/golang-rest-clean-architecture/usecase/interactor"
	"github.com/malma28/golang-rest-clean-architecture/usecase/presenter/implement"
)

type serverGorillaMux struct {
	router     *mux.Router
	httpServer *http.Server
}

func newServerGorillaMux() Server {
	server := new(serverGorillaMux)
	server.router = mux.NewRouter()

	return server
}

func (server *serverGorillaMux) Listen(host string, port int) error {
	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%v:%v", host, port),
		Handler: server.router,
	}

	if err := server.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (server *serverGorillaMux) Shutdown() error {
	return server.httpServer.Close()
}

func (server *serverGorillaMux) Setup(validator validator.Validator, db database.SQL) error {
	if err := server.setupUserController(server.router.PathPrefix("/customers").Subrouter(), validator, db); err != nil {
		return err
	}

	return nil
}

func (server *serverGorillaMux) setupUserController(router *mux.Router, validator validator.Validator, db database.SQL) error {
	userController := controllerimplement.NewCustomerController(
		validator,
		interactor.NewCustomerInteractor(
			mysql.NewCustomerRepository(db),
			implement.NewCustomerPresenter(),
		),
	)

	responseBadRequest := new(response.ResponsePayload).SetData(nil).
		SetMessage("Bad Request").SetStatus(http.StatusBadRequest).SetSuccess(false)

	if err := router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			w.WriteHeader(responseBadRequest.StatusCode)
			w.Write(responseBadRequest.Write())
			return
		}

		response := userController.CreateCustomer(r.Body, r.Header.Get("Content-Type"))
		w.WriteHeader(response.StatusCode)
		w.Write(response.Write())
	}).Methods("POST").GetError(); err != nil {
		return err
	}

	if err := router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		response := userController.GetAllCustomer()
		w.WriteHeader(response.StatusCode)
		w.Write(response.Write())
	}).Methods("GET").GetError(); err != nil {
		return err
	}

	if err := router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr, isExist := mux.Vars(r)["id"]
		if !isExist {
			w.WriteHeader(responseBadRequest.StatusCode)
			w.Write(responseBadRequest.Write())
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			w.WriteHeader(responseBadRequest.StatusCode)
			w.Write(responseBadRequest.Write())
			return
		}

		response := userController.GetCustomerById(id)
		w.WriteHeader(response.StatusCode)
		w.Write(response.Write())
	}).Methods("GET").GetError(); err != nil {
		return err
	}

	if err := router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr, isExist := mux.Vars(r)["id"]
		if !isExist {
			w.WriteHeader(responseBadRequest.StatusCode)
			w.Write(responseBadRequest.Write())
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			w.WriteHeader(responseBadRequest.StatusCode)
			w.Write(responseBadRequest.Write())
			return
		}

		response := userController.UpdateCustomerById(id, r.Body, r.Header.Get("Content-Type"))
		w.WriteHeader(response.StatusCode)
		w.Write(response.Write())
	}).Methods("PUT").GetError(); err != nil {
		return err
	}

	if err := router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr, isExist := mux.Vars(r)["id"]
		if !isExist {
			w.WriteHeader(responseBadRequest.StatusCode)
			w.Write(responseBadRequest.Write())
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			w.WriteHeader(responseBadRequest.StatusCode)
			w.Write(responseBadRequest.Write())
			return
		}

		response := userController.DeleteCustomerById(id)
		w.WriteHeader(response.StatusCode)
		w.Write(response.Write())
	}).Methods("DELETE").GetError(); err != nil {
		return err
	}

	return nil
}
