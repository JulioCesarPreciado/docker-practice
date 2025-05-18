package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

// ConnectRedis initializes and returns a Redis client using environment variables.
// It verifies the connection by pinging the Redis server.
// Returns the client and any connection error encountered.
func ConnectRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // default DB
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
