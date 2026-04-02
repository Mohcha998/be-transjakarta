package config

import "os"

type Config struct {
	DB string
}

func Load() Config {
	return Config{
		DB: os.Getenv("DB_DSN"),
	}
}
