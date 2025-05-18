// Package websocket provides functionality to upgrade HTTP connections to WebSocket
// and handle basic message communication between the client and server.
package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"example-chat/internal/flow"
)

var hub = NewHub()

func init() {
	go hub.Run()
}

// upgrader configures the WebSocket upgrade behavior, including origin checking.
// In this implementation, all origins are accepted (should be restricted in production).
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // ⚠️ Permite todas las conexiones — mejorar en producción
	},
}

// HandleWebSocket upgrades an incoming HTTP request to a WebSocket connection,
// registers the client with the hub, and starts read/write goroutines.
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("❌ Error al hacer upgrade a WebSocket:", err)
		return
	}

	client := NewClient(conn)
	hub.Register <- client

	go handleRead(client)
	go handleWrite(client)
}

// handleRead continuously reads messages from the WebSocket connection and sends them to the hub.
func handleRead(c *Client) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("👋 Cliente cerró la conexión:", err)
			break
		}
		response := flow.ProcessMessage(msg, c.State)
		c.Send <- response
	}
}

// handleWrite continuously writes messages from the Send channel to the WebSocket connection.
func handleWrite(c *Client) {
	defer c.Conn.Close()
	for msg := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("❌ Error al escribir mensaje:", err)
			break
		}
	}
}
