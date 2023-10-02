package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigOptions struct {
	ENV                 string
	SERVER              string
	PORT                string
	API_VERSION         string
	EXPIRED_CONNECTIONS uint
	WORKING_DIR         string

	DATABASE_HOST     string
	DATABASE_USER     string
	DATABASE_PORT     string
	DATABASE_PASSWORD string
	DATABASE_DB_NAME  string
	DATABASE_SSL_MODE string
	DATABASE_TIMEZONE string
}

var Config ConfigOptions = ConfigOptions{}

func LoadConfig() error {
	// Memuat variabel lingkungan dari berkas .env
	if err := godotenv.Load(); err != nil {
		return err
	}

	Config.ENV = os.Getenv("GO_ENV")
	conf, errLoadConf := LoadEnvConfig()
	if errLoadConf != nil {
		return errLoadConf
	}
	Config = conf

	return nil
}
