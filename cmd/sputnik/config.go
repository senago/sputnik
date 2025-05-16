package main

import (
	"os"

	"github.com/senago/sputnik/internal/ioc"
)

func GetConfig() ioc.Config {
	return ioc.Config{
		DSN: os.Getenv("PG_DSN"),
	}
}
