// Package main initializes and starts the WebSocket server for the Example chat system.
// It loads environment variables, configures routes, and listens for incoming HTTP connections.
package main

import (
	"log"
	"net/http"
	"os"

	// Required for handling WebSocket upgrades and communication
	"example-chat/internal/websocket"

	"github.com/joho/godotenv"
)

// main is the entry point of the WebSocket server.
// It loads environment variables from a .env file (if present),
// sets up the WebSocket route, and starts the HTTP server on the specified port.
func main() {
	// 📦 Cargar variables del .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se pudo cargar el .env, se usarán valores por defecto")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Register the WebSocket route for handling real-time client communication
	// 🌐 Ruta para WebSocket
	http.HandleFunc("/ws", websocket.HandleWebSocket)

	// 🚀 Iniciar servidor
	log.Printf("🚀 Servidor escuchando en http://localhost:%s/ws\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("❌ Error al iniciar servidor:", err)
	}
}
