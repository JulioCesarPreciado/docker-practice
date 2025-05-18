package main

import (
	// api handles the route definitions and HTTP handler wiring
	"example-api/api"
	"example-api/pkg/logger"

	// config provides functions for environment configuration and database connection
	"example-api/pkg/config"

	"fmt"
	"log"
	"net/http"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

func main() {
	logger.Setup("server.log")
	// Initialize the database connection using environment variables
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the Redis client using environment variables
	redisClient, err := config.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()

	fmt.Println("🚀 API running on http://localhost:8080")

	// Start the HTTP server with the registered routes
	http.ListenAndServe(":8080", api.RegisterRoutes(db, redisClient))
}
