package wsx

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// clientsMappingTokens save the mapping between the client and the token.
	clientsMappingTokens map[string]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	upgrader *websocket.Upgrader
}

func NewHub(upgrader *websocket.Upgrader) *Hub {
	return &Hub{
		broadcast:            make(chan []byte),
		register:             make(chan *Client),
		unregister:           make(chan *Client),
		clients:              make(map[*Client]bool),
		clientsMappingTokens: make(map[string]*Client),
		upgrader:             upgrader,
	}
}

// UpgradeToWs upgrades the HTTP connection to a WebSocket connection.
func (h *Hub) UpgradeToWs(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	return conn, err
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.clientsMappingTokens[client.token] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				h.clear(client)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					h.clear(client)
				}
			}
		}
	}
}

func (h *Hub) AddClient(conn *websocket.Conn, token string) *Client {
	client := NewClient(conn, h, token)
	h.register <- client
	return client
}

func (h *Hub) clear(client *Client) {
	delete(h.clients, client)
	delete(h.clientsMappingTokens, client.token)
	close(client.send)
}
