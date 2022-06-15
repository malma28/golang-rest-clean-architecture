package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/malma28/golang-rest-clean-architecture/adapter/database"
)

type mysqlTxHandler struct {
	tx *sql.Tx
}

func (txHandler *mysqlTxHandler) Exec(ctx context.Context, query string, args ...any) (database.Result, error) {
	return txHandler.tx.ExecContext(ctx, query, args...)
}

func (txHandler *mysqlTxHandler) Query(ctx context.Context, query string, args ...any) (database.Rows, error) {
	return txHandler.tx.QueryContext(ctx, query, args...)
}

func (txHandler *mysqlTxHandler) QueryRow(ctx context.Context, query string, args ...any) (database.Row, error) {
	row := txHandler.tx.QueryRowContext(ctx, query, args...)
	return row, row.Err()
}

func (txHandler *mysqlTxHandler) Commit() error {
	return txHandler.tx.Commit()
}

func (txHandler *mysqlTxHandler) Rollback() error {
	return txHandler.tx.Rollback()
}

type mysqlHandler struct {
	db *sql.DB
}

func newMysqlHandler(config SQLConfig) database.SQL {
	var err error

	mysqlConfig := &mysql.Config{
		User:                 config.Username,
		Passwd:               config.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", config.Host, config.Port),
		DBName:               config.DatabaseName,
		AllowNativePasswords: config.Options.AllowNativePassword,
		MultiStatements:      config.Options.MultiStatements,
		ParseTime:            config.Options.ParseTimes,
	}

	handler := new(mysqlHandler)
	handler.db, err = sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	return handler
}

func (handler *mysqlHandler) Exec(ctx context.Context, query string, args ...any) (database.Result, error) {
	return handler.db.ExecContext(ctx, query, args)
}

func (handler *mysqlHandler) Query(ctx context.Context, query string, args ...any) (database.Rows, error) {
	return handler.db.QueryContext(ctx, query, args...)
}

func (handler *mysqlHandler) QueryRow(ctx context.Context, query string, args ...any) (database.Row, error) {
	row := handler.db.QueryRowContext(ctx, query, args...)
	return row, row.Err()
}

func (handler *mysqlHandler) Tx(ctx context.Context) (database.Tx, error) {
	tx, err := handler.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &mysqlTxHandler{tx}, nil
}

func (handler *mysqlHandler) Close() error {
	return handler.db.Close()
}
