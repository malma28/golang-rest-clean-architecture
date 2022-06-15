package server

import (
	"github.com/malma28/golang-rest-clean-architecture/adapter/database"
	"github.com/malma28/golang-rest-clean-architecture/adapter/validator"
)

type Server interface {
	Listen(host string, port int) error
	Shutdown() error
	Setup(validator validator.Validator, db database.SQL) error
}

type ServerType int

const (
	ServerGorillaMux ServerType = iota
)

func NewServer(serverType ServerType) Server {
	switch serverType {
	case ServerGorillaMux:
		return newServerGorillaMux()
	}
	return nil
}
