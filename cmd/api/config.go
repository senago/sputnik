package main

import (
	"os"
)

func GetDSN() string {
	return os.Getenv("PG_DSN")
}
