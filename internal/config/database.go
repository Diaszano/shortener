// Package config provides utilities for configuring and establishing connections
// to external resources, including the database.
package config

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDatabase establishes a connection to the database using a connection pool.
//
// Parameters:
// - ctx: A context that controls the connection process, allowing for cancellation or timeout.
//
// Returns:
// - *pgxpool.Pool: A connection pool for interacting with the database.
// - error: An error if the connection fails or if the database cannot be reached.
func ConnectDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	pool, err := pgxpool.New(ctx, Env.Database.Dsn())
	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
