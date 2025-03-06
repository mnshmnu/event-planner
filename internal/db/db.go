package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func Init() (conn *pgxpool.Pool, err error) {

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("CONNECTION_STRING"))
	if err != nil {
		zap.S().Errorf("failed to connect to database: %v", err)
		return nil, err
	}

	err = InitTables(dbpool)
	if err != nil {
		zap.S().Errorf("failed to initialize tables: %v", err)
		return nil, err
	}

	return dbpool, nil
}
