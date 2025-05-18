// Package websocket manages WebSocket clients and message flow.
package websocket

import (
	"github.com/gorilla/websocket"
)

// Client represents a WebSocket connection to a single user.
// It includes the connection, a channel for outbound messages,
// an optional identifier, and a conversation state used to track
// the user's responses in a chat-like form flow.
type Client struct {
	Conn  *websocket.Conn   // The WebSocket connection.
	Send  chan []byte       // Outbound messages to be sent to this client.
	ID    string            // Optional client identifier.
	State map[string]string // Stores form-like user input step-by-step.
}

// NewClient creates and initializes a new WebSocket client instance.
// It sets up the connection, the message channel, and initializes
// the conversation state storage map.
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn:  conn,
		Send:  make(chan []byte),
		State: make(map[string]string),
	}
}
