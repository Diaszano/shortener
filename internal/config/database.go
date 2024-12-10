// Package config provides utilities for configuring and establishing connections
// to external resources, including the database.
package config

import (
	"context"

	"github.com/agoda-com/opentelemetry-go/otelzap"
	"go.uber.org/zap"

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
	otelzap.Ctx(ctx).Info("Starting connection to the database")

	select {
	case <-ctx.Done():
		otelzap.Ctx(ctx).Warn("Context canceled before connecting to the database")
		return nil, ctx.Err()

	default:
		pool, err := pgxpool.New(ctx, Env.Database.Dsn())
		if err != nil {
			otelzap.Ctx(ctx).Error("Error creating new database connection pool", zap.Error(err))
			return nil, err
		}

		otelzap.Ctx(ctx).Info("Database connection pool created successfully")

		err = pool.Ping(ctx)
		if err != nil {
			otelzap.Ctx(ctx).Error("Error pinging the database", zap.Error(err))
			pool.Close()
			return nil, err
		}

		otelzap.Ctx(ctx).Info("Successfully connected to the database")
		return pool, nil
	}
}
