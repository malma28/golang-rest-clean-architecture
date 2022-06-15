package database

import "github.com/malma28/golang-rest-clean-architecture/adapter/database"

type SQLDatabaseType int

const (
	SQLDatabaseMySQL SQLDatabaseType = iota
)

func NewSQLDatabase(sqlDatabaseType SQLDatabaseType, sqlConig SQLConfig) database.SQL {
	switch sqlDatabaseType {
	case SQLDatabaseMySQL:
		return newMysqlHandler(sqlConig)
	}
	return nil
}
