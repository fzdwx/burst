package wsx

import (
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"github.com/gorilla/websocket"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// onBinary is called when a binaryMessage is received on the websocket.
	onBinary func([]byte)

	// onText is called when a textMessage is received on the websocket.
	onText func(string)

	// onError is called when the connection has an error.
	// return true to close the connection.
	onError func(error) bool

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *Client) OnBinary(f func(data []byte)) {
	c.onBinary = f
}

func (c *Client) OnText(f func(data string)) {
	c.onText = f
}

func (c *Client) OnError(f func(err error) bool) {
	c.onError = f
}

func (c Client) Write(binary []byte) {
	c.send <- binary
}

func NewClient(conn *websocket.Conn, h *Hub) *Client {
	client := &Client{
		hub:  h,
		conn: conn,
		send: make(chan []byte, bufSize),

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

		onError: func(err error) bool {
			logx.Errorw("onError", logx.LogField{
				Key:   "err",
				Value: err,
			})
			return false
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
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		msgType, message, err := c.conn.ReadMessage()
		if err != nil {
			if c.onError(err) {
				break
			}
		}

		if msgType == websocket.BinaryMessage {
			c.onBinary(message)
		} else if msgType == websocket.TextMessage {
			c.onText(string(message))
		} else if msgType == websocket.CloseMessage {
			break
		} else {
			logx.Info("onSupport message:", msgType, string(message))
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
		c.conn.Close()
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
