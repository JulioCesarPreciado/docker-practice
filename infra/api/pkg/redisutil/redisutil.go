package redisutil

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// SetJSON stores any struct as JSON in Redis with TTL.
func SetJSON[T any](client *redis.Client, key string, value T, ttl time.Duration) error {
	ctx := context.Background()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, data, ttl).Err()
}

// GetJSON retrieves a value from Redis and deserializes it into the provided pointer.
func GetJSON[T any](client *redis.Client, key string) (T, bool, error) {
	var result T
	ctx := context.Background()

	data, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return result, false, nil
	} else if err != nil {
		return result, false, err
	}

	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return result, false, err
	}

	return result, true, nil
}
