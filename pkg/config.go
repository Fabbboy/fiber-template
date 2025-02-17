package pkg

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DEF_DATABASE_URL = "mysql://root:password@tcp(localhost:3306)/fiber_template"

type Config struct {
	Host          string // localhost:1221
	DatabaseUrl   string // mysql://root:password@localhost:3306/fiber_template
	DbConnRetries int    // 5
}

func LookupOrDefault(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}

func LookupOrDefaultInt(key string, def int) int {
	if value, ok := os.LookupEnv(key); ok {

		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		return def
	}
	return def
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		Host:          LookupOrDefault("HOST", "localhost:1221"),
		DatabaseUrl:   LookupOrDefault("DATABASE_URL", DEF_DATABASE_URL),
		DbConnRetries: LookupOrDefaultInt("DB_CONN_RETRIES", 5),
	}
}
