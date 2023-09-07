package config

import (
	"fmt"
	"go-ws-server/src/middleware/converts"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var LoadEnvConfig = func(env string) (ConfigOptions, error) {
	config := ConfigOptions{}
	if err := godotenv.Load(); err != nil {
		return config, err
	}
	ENV := strings.ToUpper(env)
	config.ENV = env
	config.SERVER = os.Getenv(fmt.Sprintf("%v_SERVER", ENV))
	config.PORT = os.Getenv(fmt.Sprintf("%v_PORT", ENV))
	config.API_VERSION = os.Getenv("API_VERSION")
	config.EXPIRED_CONNECTIONS = uint(converts.StringToInt(os.Getenv("EXPIRED_CONNECTIONS")))
	config.WORKING_DIR = os.Getenv(fmt.Sprintf("%v_WORKING_DIR", ENV))

	config.DATABASE_HOST = os.Getenv(fmt.Sprintf("%v_DATABASE_HOST", ENV))
	config.DATABASE_PORT = os.Getenv(fmt.Sprintf("%v_DATABASE_PORT", ENV))
	config.DATABASE_USER = os.Getenv(fmt.Sprintf("%v_DATABASE_USERNAME", ENV))
	config.DATABASE_PASSWORD = os.Getenv(fmt.Sprintf("%v_DATABASE_PASSWORD", ENV))
	config.DATABASE_DB_NAME = os.Getenv(fmt.Sprintf("%v_DATABASE_DBNAME", ENV))
	config.DATABASE_SSL_MODE = os.Getenv(fmt.Sprintf("%v_DATABASE_SSL_MODE", ENV))
	config.DATABASE_TIMEZONE = os.Getenv(fmt.Sprintf("%v_DATABASE_TIMEZONE", ENV))

	return config, nil
}
