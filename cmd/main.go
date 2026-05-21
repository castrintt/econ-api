package main

import (
	"os"

	"github.com/castrintt/econ-api/cmd/api"
	"github.com/castrintt/econ-api/cmd/env"
	"github.com/castrintt/econ-api/cmd/logger"
)

func main() {
	env.InitializeEnvironmentVariables()
	slog := logger.InitializeLogger()
	if err := api.InitializeApplication(); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
