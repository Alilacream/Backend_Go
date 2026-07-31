package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Secret string
}

func envAll() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Couldn't Load Env Vars \n")
		return nil
	}
	return &Config{
		Secret: os.Getenv("SECRET"),
	}
}

func GetEnv(key string) any {
	allKeys := envAll()
	switch key {
	case "SECRET":
		return allKeys.Secret
	default:
		return "No key"
	}
}
