// Package config provides utilities for configuring and connecting to external services
// like Redis cache within the application.
package config

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

// ConnectCache establishes a connection to the Redis cache using the application configuration.
// It uses a context with a timeout to limit the duration of the connection attempt.
func ConnectCache(ctx context.Context) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	default:
		options := redis.Options{
			Addr:     Env.Cache.Addr(),
			Password: Env.Cache.Pass,
			DB:       int(Env.Cache.DB),
		}

		client := redis.NewClient(&options)

		_, err := client.Ping().Result()
		if err != nil {
			return nil, err
		}

		return client, nil
	}
}
