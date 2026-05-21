package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST          string
	DB_PORT          string
	DB_USER          string
	DB_PASSWORD      string
	DB_NAME          string
	APPLICATION_PORT string
)

func InitializeEnvironmentVariables() {
	_ = godotenv.Load()

	DB_HOST = getenvOrDefault("DB_HOST", "localhost")
	DB_PORT = getenvOrDefault("DB_PORT", "5432")
	DB_USER = getenvOrDefault("DB_USER", "postgres")
	DB_PASSWORD = getenvOrDefault("DB_PASSWORD", "")
	DB_NAME = getenvOrDefault("DB_NAME", "postgres")
	APPLICATION_PORT = getenvOrDefault("APPLICATION_PORT", ":8080")
}

func getenvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return defaultValue
}
