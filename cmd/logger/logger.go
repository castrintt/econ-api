package logger

import (
	"log/slog"
	"os"
)

func InitializeLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}