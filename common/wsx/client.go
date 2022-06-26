package wsx

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	token string

	// The websocket connection.
	conn *websocket.Conn

	// onBinary is called when a binaryMessage is received on the websocket.
	onBinary func([]byte)

	// onText is called when a textMessage is received on the websocket.
	onText func(string)

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *Client) OnBinary(f func(data []byte)) {
	if f == nil {
		return
	}
	c.onBinary = f
}

func (c *Client) OnText(f func(data string)) {
	if f == nil {
		return
	}
	c.onText = f
}

func (c Client) Write(binary []byte) {
	if binary == nil {
		return
	}
	c.send <- binary
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c Client) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func NewClient(conn *websocket.Conn, h *Hub, token string) *Client {
	client := &Client{
		token: token,
		hub:   h,
		conn:  conn,
		send:  make(chan []byte, bufSize),

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

	return client
}

// ReadPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
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

// WritePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
