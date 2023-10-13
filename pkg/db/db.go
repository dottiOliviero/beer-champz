package db

import (
	"beerchampz/pkg/config"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// GetDb :
func GetDb(conf *config.Config) *pgxpool.Pool {
	// dsn := "host=localhost user=test password=test dbname=test port=5432 sslmode=disable"
	connString := conf.GetDBConnStr()
	logger, _ := zap.NewProduction()
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, connString)
	if err != nil {
		logger.Fatal("Unable to connect to database", zap.Error(err))
	}

	if err = conn.Ping(ctx); err != nil {
		logger.Fatal("Error Pinging database", zap.Error(err))
	}

	return conn
}
