package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	GOOSE_DBSTRING      string
	GOOSE_DRIVER        string
	GOOSE_MIGRATION_DIR string
	APPLICATION_PORT    string
)

func InitializeEnvironmentVariables() {
	_ = godotenv.Load()
	GOOSE_DBSTRING = getenvOrDefault("GOOSE_DBSTRING", "")
	GOOSE_DRIVER = getenvOrDefault("GOOSE_DRIVER", "")
	GOOSE_MIGRATION_DIR = getenvOrDefault("GOOSE_MIGRATION_DIR", "")
	APPLICATION_PORT = getenvOrDefault("APPLICATION_PORT", ":8080")
}

func getenvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return defaultValue
}
