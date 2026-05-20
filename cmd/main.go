package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config := configuration{
		port: getenvOrDefault("APPLICATION_PORT", ":8080"),
		database: databaseConfiguration{
			host:         getenvOrDefault("DB_HOST", "localhost"),
			port:         getenvOrDefault("DB_PORT", "5432"),
			user:         getenvOrDefault("DB_USER", "postgres"),
			password:     os.Getenv("DB_PASSWORD"),
			databaseName: getenvOrDefault("DB_NAME", "postgres"),
		},
	}

	api := &application{
		configuration: config,
	}
	handler := api.mount()
	api.run(handler)
}
