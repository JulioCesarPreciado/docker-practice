package year

// Package year provides HTTP handlers and business logic for vehicle year catalog operations.

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// GetYearsHandler handles HTTP GET requests to retrieve vehicle years filtered by model_id.
// It requires a model_id query parameter, and returns a 400 Bad Request if it's missing.
func GetYearsHandler(db *sqlx.DB, w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {

	modelID := r.URL.Query().Get("model_id")
	if modelID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := NewRepository(db, redisClient)

	years, err := repo.GetAllYears(modelID)
	if err != nil {
		http.Error(w, "Failed to fetch years", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(years)
}
