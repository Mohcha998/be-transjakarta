package config

import "os"

type Config struct {
	DB     string
	Rabbit string
}

func Load() Config {
	return Config{
		DB:     os.Getenv("DB_DSN"),
		Rabbit: os.Getenv("RABBIT_URL"),
	}
}
