package svc

import (
	"github.com/fzdwx/burst/internal/server"
	"github.com/gorilla/websocket"
	"net/http"
)

type (
	ServiceContext struct {
		Config     server.Config
		WsUpgrader websocket.Upgrader
	}
)

func NewServiceContext(c server.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		WsUpgrader: newWsUpgrader(c),
	}
}

// newWsUpgrader new websocket.Upgrader instance
func newWsUpgrader(c server.Config) websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
