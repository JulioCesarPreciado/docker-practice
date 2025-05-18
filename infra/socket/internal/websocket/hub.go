// Package websocket implements the hub that manages all connected WebSocket clients.
package websocket

import (
	"encoding/json"
	"log"

	"example-chat/internal/flow"
)

// Hub maintains the set of active clients and manages their registration and unregistration.
// The broadcast mechanism is currently unused but retained for potential system-wide messaging.
type Hub struct {
	Clients    map[*Client]bool // Registered clients connected to the hub.
	Broadcast  chan []byte      // Deprecated: Used for broadcasting to all clients (no longer used).
	Register   chan *Client     // Channel for registering new clients with the hub.
	Unregister chan *Client     // Channel for unregistering disconnected clients.
}

// NewHub initializes and returns a new Hub instance with empty client storage
// and active channels for client management.
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run starts the hub's main event loop. It listens for new client registrations,
// unregistrations, and (deprecated) broadcast messages. Each client is tracked
// in the Clients map and removed upon disconnection.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			log.Println("✅ Cliente registrado")
			// Send a welcome message to the new client, using the first flow step definition.
			if len(flow.StepDefinitions) > 0 {
				first := flow.StepDefinitions[0].Name
				if cfg, ok := flow.GetStepConfig(first); ok {
					if jsonCfg, err := json.Marshal(cfg); err == nil {
						client.Send <- jsonCfg
					}
				}
			}

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				log.Println("👋 Cliente desconectado")
			}

		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
