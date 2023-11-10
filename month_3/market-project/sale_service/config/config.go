package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Limit int
	Page  int

	Environment string // debug, test, release

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	Port string

	PostgresMaxConnections int32

	DefaultOffset int
	DefaultLimit  int

	// redis configuration
	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDatabase int
}

const (
	TokenExpireTime = 24 * time.Hour
	JWTSecretKey    = "MySecretKey"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"

	TimeExpiredAt = time.Hour * 720
)

// Load ...
func Load() Config {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "2511"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "blog"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	// redis configuration
	config.RedisHost = cast.ToString(getOrReturnDefaultValue("REDIS_HOST", "localhost"))
	config.RedisPort = cast.ToInt(getOrReturnDefaultValue("REDIS_PORT", 6379))
	config.RedisPassword = cast.ToString(getOrReturnDefaultValue("REDIS_PASSWORD", ""))
	config.RedisDatabase = cast.ToInt(getOrReturnDefaultValue("REDIS_DATABASE", 0))

	return config

}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
