package wsx

import (
	"github.com/gorilla/websocket"
	"net/http"
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
