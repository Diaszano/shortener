// Package config provides utilities for configuring and connecting to external services
// like Redis cache within the application.
package config

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

// ConnectCache establishes a connection to the Redis cache using the application configuration.
// It uses a context with a timeout to limit the duration of the connection attempt.
//
// Parameters:
//   - ctx: The context to control timeout and cancellation.
//
// Returns:
//   - *redis.Client: A pointer to the initialized Redis client.
//   - error: An error if the connection fails.
func ConnectCache(ctx context.Context) (*redis.Client, error) {
	log := GetLogger()
	log.Info("Starting connection to the cache")

	select {
	case <-ctx.Done():
		log.Warn("Context canceled before connecting to the cache")
		return nil, ctx.Err()

	default:
		options := redis.Options{
			Addr:     Env.Cache.Addr(),
			Password: Env.Cache.Pass,
			DB:       int(Env.Cache.DB),
		}

		log.Debug("Redis options configured", zap.String("Addr", options.Addr), zap.Int("DB", options.DB))

		client := redis.NewClient(&options)

		_, err := client.Ping().Result()
		if err != nil {
			log.Error("Error connecting to the cache", zap.Error(err))
			_ = client.Close()
			return nil, err
		}

		log.Info("Successfully connected to the cache")
		return client, nil
	}
}
