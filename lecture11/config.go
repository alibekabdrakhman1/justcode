package main

import (
	"time"
)

type Config struct {
	HttpServer HttpServerConfig `mapstructure:"HttpServer"`
	Postgres   PostgresConfig   `mapstructure:"Postgres"`
}

type HttpServerConfig struct {
	Port            int           `mapstructure:"Port"`
	ShutdownTimeout time.Duration `mapstructure:"ShutdownTimeout"`
}
type PostgresConfig struct {
	Port   int    `mapstructure:"Port"`
	DBName string `mapstructure:"DBName"`
}
