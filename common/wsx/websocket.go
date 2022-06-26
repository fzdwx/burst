package wsx

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

// Websocket is a middleman between the websocket connection and the hub.
type Websocket struct {
	hub *Hub

	token string

	// The websocket connection.
	conn *websocket.Conn

	// onBinary is called when a binaryMessage is received on the websocket.
	onBinary func([]byte)

	// onText is called when a textMessage is received on the websocket.
	onText func(string)
}

func (c *Websocket) OnBinary(f func(data []byte)) {
	if f == nil {
		return
	}
	c.onBinary = f
}

func (c *Websocket) OnText(f func(data string)) {
	if f == nil {
		return
	}
	c.onText = f
}

func (c Websocket) Write(binary []byte) error {
	return c.conn.WriteMessage(websocket.BinaryMessage, binary)
}

func (c Websocket) WriteStr(s string) error {
	return c.conn.WriteMessage(websocket.TextMessage, []byte(s))
}

// Close closes the websocket connection.
func (c *Websocket) Close() error {
	return c.conn.Close()
}

// RemoteAddr returns the remote address of the peer.
func (c Websocket) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func NewClient(conn *websocket.Conn, h *Hub, token string) *Websocket {
	client := &Websocket{
		token: token,
		hub:   h,
		conn:  conn,

		onBinary: func(bytes []byte) {
			logx.Sloww("onBinary", logx.LogField{
				Key:   "msg",
				Value: bytes,
			})
		},

		onText: func(s string) {
			logx.Sloww("onText", logx.LogField{
				Key:   "msg",
				Value: s,
			})
		},
	}
	conn.SetCloseHandler(func(code int, text string) error {
		logx.Sloww("websocket close", logx.LogField{
			Key:   "code",
			Value: code,
		})

		fmt.Println("close")
		return nil
	})
	return client
}

// React pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Websocket) React() {
	defer func() {
		c.hub.unregister <- c
		c.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		msgType, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		if msgType == websocket.BinaryMessage {
			c.onBinary(message)
		} else if msgType == websocket.TextMessage {
			c.onText(string(message))
		} else if msgType == websocket.CloseMessage {
			break
		} else {
			logx.Info("unSupport message:", msgType, string(message))
		}
	}
}
