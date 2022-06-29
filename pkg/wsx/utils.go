package wsx

import "github.com/gorilla/websocket"

var ping = func(conn *websocket.Conn) error {
	return conn.WriteMessage(websocket.PingMessage, nil)
}

// Ping websocket conn
func Ping(conn *websocket.Conn) error {
	return ping(conn)
}