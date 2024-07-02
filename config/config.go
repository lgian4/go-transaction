package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

type Config struct {
	Host        string `validate:"required"`
	Port        int64  `validate:"required,numeric"`
	DbName      string `validate:"required"`
	GinMode     string `validate:"required"`
	LogFilename string `validate:"required"`
}

func Load(filename string) (Config, error) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config := Config{}
	host := os.Getenv("HOST")
	port, err := strconv.ParseInt(os.Getenv("Port"), 10, 64)
	if err != nil {
		return config, errors.Join(errors.New("failed to get port env"), err)
	}
	dbName := os.Getenv("DB_NAME")
	ginMode := os.Getenv("GIN_MODE")
	LogFilename := os.Getenv("LOG_FILENAME")

	config = Config{
		Host:        host,
		Port:        port,
		DbName:      dbName,
		GinMode:     ginMode,
		LogFilename: LogFilename,
	}
	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {

		return config, errors.Join(errors.New("failed to validate config"), err)
	}
	return config, nil

}
