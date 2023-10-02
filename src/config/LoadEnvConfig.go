package config

import (
	"go-ws-server/src/middleware/converts"
	"os"

	"github.com/joho/godotenv"
)

var LoadEnvConfig = func() (ConfigOptions, error) {
	config := ConfigOptions{}
	if err := godotenv.Load(); err != nil {
		return config, err
	}
	config.SERVER = os.Getenv("SERVER")
	config.PORT = os.Getenv("PORT")
	config.API_VERSION = os.Getenv("API_VERSION")
	config.EXPIRED_CONNECTIONS = uint(converts.StringToInt(os.Getenv("EXPIRED_CONNECTIONS")))
	config.WORKING_DIR = os.Getenv("WORKING_DIR")

	config.DATABASE_HOST = os.Getenv("DATABASE_HOST")
	config.DATABASE_PORT = os.Getenv("DATABASE_PORT")
	config.DATABASE_USER = os.Getenv("DATABASE_USERNAME")
	config.DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	config.DATABASE_DB_NAME = os.Getenv("DATABASE_DBNAME")
	config.DATABASE_SSL_MODE = os.Getenv("DATABASE_SSL_MODE")
	config.DATABASE_TIMEZONE = os.Getenv("DATABASE_TIMEZONE")

	return config, nil
}
