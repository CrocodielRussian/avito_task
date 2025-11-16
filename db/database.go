package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context) (*pgxpool.Pool, error) {
	dbURL := "postgres://postgres:postgres@localhost:5435/pr_database?sslmode=disable"
	return pgxpool.New(ctx, dbURL)
}
