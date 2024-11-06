package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oidentity/auth-server/internal/config"
	"github.com/oidentity/auth-server/internal/logger"
)

func ConnectPostgres() (*pgxpool.Pool, error) {

	dsn, err := buildDSN()
	if err != nil {
		return nil, err
	}

	config, err := parseConfig(dsn)
	if err != nil {
		return nil, err
	}

	dbpool, err := createPool(context.Background(), config)
	if err != nil {
		return nil, err
	}

	err = testConnection(dbpool)
	if err != nil {
		dbpool.Close()
		return nil, err
	}

	logger.GetLogger().Info("Connected to PostgreSQL successfully.")
	return dbpool, nil
}

func buildDSN() (string, error) {
	config := config.LoadConfig()
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName,
	)
	return dsn, nil
}

func parseConfig(dsn string) (*pgxpool.Config, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse PostgreSQL DSN: %v", err)
	}
	return config, nil
}

func createPool(ctx context.Context, config *pgxpool.Config) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create PostgreSQL connection pool: %v", err)
	}
	return dbpool, nil
}

func testConnection(dbpool *pgxpool.Pool) error {
	err := dbpool.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("unable to connect to PostgreSQL: %v", err)
	}
	return nil
}
