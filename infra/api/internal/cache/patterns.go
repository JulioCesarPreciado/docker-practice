// Package cache defines Redis key patterns for each supported catalog type.
// These patterns are used to match and invalidate cache entries by category.
package cache

// CatalogPatterns maps catalog names to their corresponding Redis key patterns.
var CatalogPatterns = map[string]string{
	"brands":   "brands",
	"models":   "model:*",
	"years":    "year:*",
	"versions": "version:*",
}
