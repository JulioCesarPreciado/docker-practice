package version

// Package version provides HTTP handlers and business logic for vehicle version catalog operations.

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// GetVersionsHandler handles HTTP GET requests to retrieve vehicle versions filtered by year_id.
// It requires a year_id query parameter, and returns a 400 Bad Request if it's missing.
func GetVersionsHandler(db *sqlx.DB, w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {

	yearID := r.URL.Query().Get("year_id")
	if yearID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := NewRepository(db, redisClient)

	versions, err := repo.GetAllVersions(yearID)
	if err != nil {
		http.Error(w, "Failed to fetch versions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(versions)
}
