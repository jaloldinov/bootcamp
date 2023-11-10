package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	BranchServiceHost string
	BranchServicePort int

	LogLevel string
	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	// if err := godotenv.Load(); err != nil {
	// 	fmt.Println("No .env file found")
	// }

	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	config.BranchServiceHost = cast.ToString(getOrReturnDefault("BRANCH_SERVICE_HOST", "localhost"))
	config.BranchServicePort = cast.ToInt(getOrReturnDefault("BRANCH_SERVICE_PORT", 50051))

	// config.ProfessionServiceHost = cast.ToString(getOrReturnDefault("Profession_SERVICE_HOST", "localhost"))
	// config.ProfessionServicePort = cast.ToInt(getOrReturnDefault("Profession_SERVICE_PORT", 9103))
	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
