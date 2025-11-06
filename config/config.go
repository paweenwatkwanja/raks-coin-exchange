package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppHost  string
	AppPort  string
	Endpoint string
}

func LoadConfig() *Config {
	fmt.Println("Loading configs")
	return &Config{
		AppHost:  os.Getenv("APP_HOST"),
		AppPort:  os.Getenv("APP_PORT"),
		Endpoint: os.Getenv("ENDPOINT"),
	}
}
