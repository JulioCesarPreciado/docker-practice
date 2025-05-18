// Package cache provides HTTP handlers and logic for managing Redis-based cache operations,
// including cache invalidation for various vehicle catalog resources.
package cache

import (
	"fmt"
	"net/http"

	"example-api/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

// RefreshCacheHandler returns an HTTP handler that clears all Redis keys matching the pattern
// associated with a given catalog name. The catalog name is expected as a URL path variable.
// Supported catalogs are defined in CatalogPatterns. If the pattern is not found, it returns 400.
// If clearing the cache fails, it returns 500. On success, it returns 204 No Content.
func RefreshCacheHandler(redis *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		catalog := mux.Vars(r)["catalog"]

		pattern, ok := CatalogPatterns[catalog]
		if !ok {
			http.Error(w, "Invalid catalog", http.StatusBadRequest)
			return
		}

		if err := ClearCacheByPattern(redis, pattern); err != nil {
			logger.Error(fmt.Sprintf("❌ Error clearing cache for pattern '%s': %v", pattern, err))
			http.Error(w, "Failed to clear cache", http.StatusInternalServerError)
			return
		}

		logger.Info(fmt.Sprintf("✅ Cache cleared for pattern: %s", pattern))
		w.WriteHeader(http.StatusNoContent)
	}
}
