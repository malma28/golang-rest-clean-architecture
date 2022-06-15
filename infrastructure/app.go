package infrastructure

import (
	"os"

	databaseadapter "github.com/malma28/golang-rest-clean-architecture/adapter/database"
	validatoradapter "github.com/malma28/golang-rest-clean-architecture/adapter/validator"
	databaseinfra "github.com/malma28/golang-rest-clean-architecture/infrastructure/database"
	serverinfra "github.com/malma28/golang-rest-clean-architecture/infrastructure/server"
	validatorinfra "github.com/malma28/golang-rest-clean-architecture/infrastructure/validator"
)

type App struct {
	server      serverinfra.Server
	validator   validatoradapter.Validator
	sqlDatabase databaseadapter.SQL
}

func New() *App {
	return new(App)
}

func (app *App) SetServer(serverType serverinfra.ServerType) *App {
	app.server = serverinfra.NewServer(serverType)
	return app
}

func (app *App) SetValidator(validatorType validatorinfra.ValidatorType) *App {
	app.validator = validatorinfra.NewValidator(validatorType)
	return app
}

func (app *App) SetSQLDatabase(sqlDatabaseType databaseinfra.SQLDatabaseType, sqlConfig databaseinfra.SQLConfig) *App {
	app.sqlDatabase = databaseinfra.NewSQLDatabase(sqlDatabaseType, sqlConfig)
	return app
}

func (app *App) Run(host string, port int, done <-chan os.Signal) error {
	if err := app.server.Setup(app.validator, app.sqlDatabase); err != nil {
		return err
	}

	doneError := make(chan error)

	go func() {
		defer os.Exit(0)
		<-done

		if err := app.server.Shutdown(); err != nil {
			doneError <- err
		}

		if err := app.sqlDatabase.Close(); err != nil {
			doneError <- err
		}

		doneError <- nil
	}()

	if err := app.server.Listen(host, port); err != nil {
		return err
	}

	return <-doneError
}
