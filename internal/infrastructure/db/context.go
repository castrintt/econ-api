package db

import (
	"context"
	"log/slog"

	"github.com/castrintt/econ-api/cmd/env"
	"github.com/jackc/pgx/v5"
)

type DatabaseConfiguration struct {
	Port             string
	ConnectionString string
}

func connect(ctx context.Context, connection string) error {
	conn, err := pgx.Connect(ctx, connection)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	return conn.Ping(ctx)
}

func DatabaseConfigurationFromEnv() DatabaseConfiguration {
	return DatabaseConfiguration{
		Port:             env.APPLICATION_PORT,
		ConnectionString: env.GOOSE_DBSTRING,
	}
}

func InitializeDatabase(ctx context.Context, connection string) error {
	if err := connect(context.Background(), connection); err != nil {
		slog.Error("Database connection failed", "error", err)
		panic(err)
	}
	return nil
}
