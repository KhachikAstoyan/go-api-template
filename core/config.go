package core

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB_URL string `envconfig:"DB_URL" required:"true"`
}

func LoadDotenvConfig() Config {
	var config Config

	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}

	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
	}

	return config
}
