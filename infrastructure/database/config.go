package database

type SQLOptions struct {
	AllowNativePassword bool
	MultiStatements     bool
	ParseTimes          bool
}

type SQLConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
	Options      SQLOptions
}
