// Package cache provides logic to manage Redis cache operations related to vehicle catalog data.
package cache

import (
	"context"
	"fmt"

	"example-api/pkg/logger"

	"github.com/redis/go-redis/v9"
)

// ClearCacheByPattern deletes all Redis keys in Redis that match the provided pattern.
// It uses SCAN to iterate over matching keys and deletes them one by one.
// If any key fails to delete, a warning is logged. The function logs each successful deletion
// and returns an error if the scan operation itself fails.
func ClearCacheByPattern(redis *redis.Client, pattern string) error {
	ctx := context.Background()
	iter := redis.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		if err := redis.Del(ctx, iter.Val()).Err(); err != nil {
			logger.Warn(fmt.Sprintf("⚠️ Failed to delete key: %s", iter.Val()))
		} else {
			logger.Info(fmt.Sprintf("🧹 Deleted cache key: %s", iter.Val()))
		}
	}
	if err := iter.Err(); err != nil {
		return fmt.Errorf("failed to clear cache for pattern %s: %w", pattern, err)
	}
	return nil
}
