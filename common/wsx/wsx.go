package wsx

import (
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"time"
)

func NewUpgrader(config struct {
	HandshakeTimeout  int64 `json:",default=10"`
	ReadBufferSize    int   `json:",default=8192"`
	WriteBufferSize   int   `json:",default=8192"`
	EnableCompression bool
}) websocket.Upgrader {
	handshakeTimeout := time.Duration(config.HandshakeTimeout) * time.Second
	return websocket.Upgrader{
		HandshakeTimeout:  handshakeTimeout,
		ReadBufferSize:    config.ReadBufferSize,
		WriteBufferSize:   config.WriteBufferSize,
		EnableCompression: config.EnableCompression,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func Connect(url url.URL, token string, hub *Hub) (*Client, *http.Response, error) {
	conn, resp, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		return nil, resp, err
	}

	return NewClient(conn, hub, token), resp, err
}
