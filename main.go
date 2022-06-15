package main

import (
	"os"
	"os/signal"

	"github.com/malma28/golang-rest-clean-architecture/infrastructure"
	"github.com/malma28/golang-rest-clean-architecture/infrastructure/database"
	"github.com/malma28/golang-rest-clean-architecture/infrastructure/server"
	"github.com/malma28/golang-rest-clean-architecture/infrastructure/validator"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	config := FromEnv()

	mainApp := infrastructure.New().SetServer(
		server.ServerGorillaMux,
	).SetSQLDatabase(
		database.SQLDatabaseMySQL,
		database.SQLConfig{
			Username:     config.MySQL.Username,
			Password:     config.MySQL.Password,
			Host:         config.MySQL.Host,
			Port:         config.MySQL.Port,
			DatabaseName: config.MySQL.DatabaseName,
			Options: database.SQLOptions{
				AllowNativePassword: config.MySQL.AllowNativePassword,
				MultiStatements:     config.MySQL.MultiStatements,
				ParseTimes:          config.MySQL.ParseTimes,
			},
		},
	).SetValidator(
		validator.ValidatorGoPlayground,
	)

	if err := mainApp.Run(config.Server.Host, config.Server.Port, done); err != nil {
		panic(err)
	}
}
