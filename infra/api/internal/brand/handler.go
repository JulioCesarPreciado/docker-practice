package brand

// Package brand provides HTTP handlers and business logic for vehicle brand catalog operations.

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// GetBrandsHandler handles HTTP GET requests to retrieve a list of vehicle brands.
// It uses the provided sqlx.DB instance to query the database and writes the result as JSON.
// Responds with HTTP 500 if the query fails.
func GetBrandsHandler(db *sqlx.DB, w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {
	repo := NewRepository(db, redisClient)

	brands, err := repo.GetAllBrands()
	if err != nil {
		http.Error(w, "Failed to fetch brands", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brands)
}
