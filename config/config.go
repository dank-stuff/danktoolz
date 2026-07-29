package config

import (
	"log"
	"os"
)

var (
	_config = config{}
)

func init() {
	_config = config{
		Port:       getEnv("PORT"),
		ApiAddress: getEnv("API_ADDRESS"),
	}
}

type config struct {
	Port       string
	AdminPort  string
	ApiAddress string
}

// Env returns the thing's config values :)
func Env() config {
	return _config
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Fatalln("The \"" + key + "\" variable is missing.")
	}
	return value
}
