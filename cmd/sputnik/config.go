package main

import "os"

type Config struct {
	DSN string
}

func GetConfig() Config {
	return Config{
		DSN: os.Getenv("PG_DSN"),
	}
}
