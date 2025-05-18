package model

// Package model provides access to vehicle model data from the database.

import (
	"example-api/pkg/redisutil"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// Repository defines the interface for querying vehicle model data.
type Repository interface {
	GetAllModels(brandID string) ([]Model, error)
}

type repository struct {
	db    *sqlx.DB
	redis *redis.Client
}

// NewRepository creates and returns a new model.Repository using the given sqlx.DB instance.
func NewRepository(db *sqlx.DB, redis *redis.Client) Repository {
	return &repository{db: db, redis: redis}
}

// GetAllModels returns all vehicle models for a given brand from the database.
// It performs a SELECT query on the vehicle_models table filtering by brand_id.
func (r *repository) GetAllModels(brandID string) ([]Model, error) {
	cacheKey := "models:" + brandID

	// Try Redis
	if models, ok, _ := redisutil.GetJSON[[]Model](r.redis, cacheKey); ok {
		return models, nil
	}

	// DB fallback
	var models []Model
	err := r.db.Select(&models, "SELECT id, name FROM vehicle_models WHERE brand_id = $1 ORDER BY name", brandID)
	if err != nil {
		return nil, err
	}

	// Set cache
	_ = redisutil.SetJSON(r.redis, cacheKey, models, time.Hour)

	return models, nil
}
