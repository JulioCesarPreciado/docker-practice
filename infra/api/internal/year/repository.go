package year

// Package year provides access to vehicle year data from the database.

import (
	"example-api/pkg/redisutil"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// Repository defines the interface for querying vehicle year data.
type Repository interface {
	GetAllYears(modelID string) ([]Year, error)
}

type repository struct {
	db    *sqlx.DB
	redis *redis.Client
}

// NewRepository creates and returns a new year.Repository using the given sqlx.DB instance.
func NewRepository(db *sqlx.DB, redis *redis.Client) Repository {
	return &repository{db: db, redis: redis}
}

// GetAllYears returns all vehicle years for a given model from the database.
// It performs a SELECT query on the vehicle_years table filtering by model_id.
func (r *repository) GetAllYears(modelID string) ([]Year, error) {
	cacheKey := "years:" + modelID

	// Try Redis
	if years, ok, _ := redisutil.GetJSON[[]Year](r.redis, cacheKey); ok {
		return years, nil
	}

	// DB fallback
	var years []Year
	err := r.db.Select(&years, "SELECT id, name FROM vehicle_years WHERE model_id = $1 ORDER BY name DESC", modelID)
	if err != nil {
		return nil, err
	}

	// Set cache
	_ = redisutil.SetJSON(r.redis, cacheKey, years, time.Hour)

	return years, nil
}
