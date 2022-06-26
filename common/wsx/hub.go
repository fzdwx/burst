package wsx

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Websocket]bool

	// clientsMappingTokens save the mapping between the client and the token.
	clientsMappingTokens map[string]*Websocket

	// Register requests from the clients.
	register chan *Websocket

	// Unregister requests from clients.
	unregister chan *Websocket

	upgrader *websocket.Upgrader
}

func NewHub(upgrader *websocket.Upgrader) *Hub {
	return &Hub{
		register:             make(chan *Websocket),
		unregister:           make(chan *Websocket),
		clients:              make(map[*Websocket]bool),
		clientsMappingTokens: make(map[string]*Websocket),
		upgrader:             upgrader,
	}
}

// UpgradeToWs upgrades the HTTP connection to a WebSocket connection.
func (h *Hub) UpgradeToWs(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	return conn, err
}

func (h *Hub) React() {
	for {
		select {
		case client := <-h.register:
			h.add(client)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				h.remove(client)
			}
		}
	}
}

func (h *Hub) AddClient(conn *websocket.Conn, token string) *Websocket {
	client := NewClient(conn, h, token)
	h.register <- client
	return client
}

// remove a client from the hub.
func (h *Hub) remove(client *Websocket) {
	defer client.Close()
	delete(h.clients, client)
	delete(h.clientsMappingTokens, client.token)
}

// add a client to the hub.
func (h *Hub) add(client *Websocket) {
	b := h.clients[client]
	if b {
		h.remove(client)
	}
	h.clients[client] = true
	h.clientsMappingTokens[client.token] = client
}

// Get  returns the client by the token.
func (h *Hub) Get(token string) *Websocket {
	return h.clientsMappingTokens[token]
}
