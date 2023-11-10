package config

import (
	"os"
	"time"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	// catalog serivce
	CatalogServiceHost string
	CatalogServicePort int

	// order serivce
	OrderServiceHost string
	OrderServicePort int

	// user serivce
	UserServiceHost string
	UserServicePort int

	LogLevel string
	HttpPort string
}

const (
	TokenExpireTime = 24 * time.Hour
	JWTSecretKey    = "MySecretKey"
)

// Load loads environment vars and inflates Config
func Load() Config {
	// if err := godotenv.Load(); err != nil {
	// 	fmt.Println("No .env file found")
	// }

	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	config.CatalogServiceHost = cast.ToString(getOrReturnDefault("CATALOG_SERVICE_HOST", "localhost"))
	config.CatalogServicePort = cast.ToInt(getOrReturnDefault("CATALOG_SERVICE_PORT", 5001))

	config.OrderServiceHost = cast.ToString(getOrReturnDefault("ORDER_SERVICE_HOST", "localhost"))
	config.OrderServicePort = cast.ToInt(getOrReturnDefault("ORDER_SERVICE_PORT", 5002))

	config.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	config.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 5003))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
