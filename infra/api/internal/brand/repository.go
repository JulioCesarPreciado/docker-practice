package brand

// Package brand provides access to vehicle brand data from the database.

import (
	"example-api/pkg/redisutil"

	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// Repository defines the interface for querying vehicle brand data.
type Repository interface {
	GetAllBrands() ([]Brand, error)
}

type repository struct {
	db    *sqlx.DB
	redis *redis.Client
}

// NewRepository creates and returns a new brand.Repository using the given sqlx.DB and redis.Client instances.
func NewRepository(db *sqlx.DB, redis *redis.Client) Repository {
	return &repository{db: db, redis: redis}
}

// GetAllBrands returns all vehicle brands from the database or cache.

func (r *repository) GetAllBrands() ([]Brand, error) {
	const cacheKey = "brands"

	// Try Redis
	if brands, ok, _ := redisutil.GetJSON[[]Brand](r.redis, cacheKey); ok {
		return brands, nil
	}

	// DB fallback
	var brands []Brand
	err := r.db.Select(&brands, "SELECT id, name FROM vehicle_brands ORDER BY name")
	if err != nil {
		return nil, err
	}

	// Set cache
	_ = redisutil.SetJSON(r.redis, cacheKey, brands, time.Hour)

	return brands, nil
}
