package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host           string
	MongoPassword  string
	MongoUser      string
	MongoURL       string
	MongoMechanism string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return &Config{}, err
	}

	return &Config{
		Host:           os.Getenv("HOST"),
		MongoPassword:  os.Getenv("MONGODB_PASSWORD"),
		MongoUser:      os.Getenv("MONGODB_USER"),
		MongoURL:       os.Getenv("MONGODB_URL"),
		MongoMechanism: os.Getenv("MONGODB_MECHANISM"),
	}, nil
}
