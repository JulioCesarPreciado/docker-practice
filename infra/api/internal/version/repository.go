package version

// Package version provides access to vehicle version data from the database.

import (
	"example-api/pkg/redisutil"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// Repository defines the interface for querying vehicle version data.
type Repository interface {
	GetAllVersions(yearID string) ([]Version, error)
}

type repository struct {
	db    *sqlx.DB
	redis *redis.Client
}

// NewRepository creates and returns a new version.Repository using the given sqlx.DB instance.
func NewRepository(db *sqlx.DB, redis *redis.Client) Repository {
	return &repository{db: db, redis: redis}
}

// GetAllVersions returns all vehicle versions for a given year from the database.
// It performs a SELECT query on the vehicle_versions table filtering by year_id.
func (r *repository) GetAllVersions(yearID string) ([]Version, error) {
	cacheKey := "versions:" + yearID

	// Try Redis
	if versions, ok, _ := redisutil.GetJSON[[]Version](r.redis, cacheKey); ok {
		return versions, nil
	}

	// DB fallback
	var versions []Version
	err := r.db.Select(&versions, "SELECT id, name FROM vehicle_versions WHERE year_id = $1 ORDER BY name", yearID)
	if err != nil {
		return nil, err
	}

	// Set cache
	_ = redisutil.SetJSON(r.redis, cacheKey, versions, time.Hour)

	return versions, nil
}
