package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// GetDb :
func GetDb() *pgxpool.Pool {
	dsn := "host=localhost user=test password=test dbname=test port=5432 sslmode=disable"
	logger, _ := zap.NewProduction()
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Fatal("Unable to connect to database", zap.Error(err))
	}

	if err = conn.Ping(ctx); err != nil {
		logger.Fatal("Error Pinging database", zap.Error(err))
	}

	return conn
}
