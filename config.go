package main

import "github.com/caarlos0/env/v6"

type ServerConfig struct {
	Port int    `env:"PORT" envDefault:"7000"`
	Host string `env:"HOST" envDefault:"127.0.0.1"`
}

type MySQLConfig struct {
	Username            string `env:"USERNAME" envDefault:"root"`
	Password            string `env:"PASSWORD,required"`
	Host                string `env:"HOST" envDefault:"127.0.0.1"`
	Port                int    `env:"PORT" envDefault:"3306"`
	DatabaseName        string `env:"DATABASE_NAME,required"`
	AllowNativePassword bool   `env:"ALLOW_NATIVE_PASSWORD" envDefault:"true"`
	MultiStatements     bool   `env:"MULTI_STATEMENTS" envDefault:"true"`
	ParseTimes          bool   `env:"PARSE_TIMES" envDefault:"true"`
}

type Config struct {
	Server ServerConfig `envPrefix:"SERVER_"`
	MySQL  MySQLConfig  `envPrefix:"MYSQL_"`
}

func FromEnv() Config {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return config
}
