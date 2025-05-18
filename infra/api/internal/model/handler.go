package model

// Package model provides HTTP handlers and business logic for vehicle model catalog operations.

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// GetModelsHandler handles HTTP GET requests to retrieve vehicle models filtered by brand_id.
// It requires a brand_id query parameter, and returns a 400 Bad Request if it's missing.
func GetModelsHandler(db *sqlx.DB, w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {

	brandID := r.URL.Query().Get("brand_id")
	if brandID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := NewRepository(db, redisClient)

	models, err := repo.GetAllModels(brandID)
	if err != nil {
		http.Error(w, "Failed to fetch models", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models)
}
