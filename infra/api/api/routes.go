// Package api provides routing logic for the HTTP server, wiring URL paths to handlers.
package api

import (
	"example-api/internal/brand"
	"example-api/internal/cache"
	"example-api/internal/model"
	"example-api/internal/version"
	"example-api/internal/year"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// RegisterRoutes sets up all HTTP route handlers using a Gorilla Mux router.
// It receives a sqlx.DB instance for dependency injection and returns the configured router.
func RegisterRoutes(db *sqlx.DB, redisClient *redis.Client) http.Handler {
	r := mux.NewRouter()
	r.Use(CORSMiddleware)

	// Apply logging middleware to all routes to capture request details and timing.
	r.Use(LoggingMiddleware)

	// Vehicle brands route
	r.HandleFunc("/brands", func(w http.ResponseWriter, r *http.Request) {
		brand.GetBrandsHandler(db, w, r, redisClient)
	}).Methods("GET")

	// Vehicle models by brands route
	r.HandleFunc("/models", func(w http.ResponseWriter, r *http.Request) {
		model.GetModelsHandler(db, w, r, redisClient)
	}).Methods("GET")

	// Vehicle years by model route
	r.HandleFunc("/years", func(w http.ResponseWriter, r *http.Request) {
		year.GetYearsHandler(db, w, r, redisClient)
	}).Methods("GET")

	// Vehicle versions by year route
	r.HandleFunc("/versions", func(w http.ResponseWriter, r *http.Request) {
		version.GetVersionsHandler(db, w, r, redisClient)
	}).Methods("GET")

	// POST /cache/refresh/{catalog} clears Redis keys matching a catalog pattern.
	// Protected by JWT authentication middleware using a shared secret from the environment.
	r.HandleFunc("/cache/refresh/{catalog}",
		RequireJWT(os.Getenv("JWT_SECRET"), cache.RefreshCacheHandler(redisClient)),
	).Methods("POST")

	return r
}
