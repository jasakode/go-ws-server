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
	return config, nil
}
