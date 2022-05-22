package cs

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type (
	Client struct {
		conn     *websocket.Conn
		token    string
		onText   OnText
		onBinary OnBinary
	}

	OnText func(string, Client)

	OnBinary func([]byte, Client)
)

func Connect(url url.URL) (*Client, error) {
	log.Printf("start connecting to %s", url.String())
	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		return nil, err
	}

	return &Client{
		conn:  c,
		token: url.Query().Get("token"),
		onText: func(s string, c Client) {
			log.Debugf("onText:%s", s)
		},
		onBinary: func(bytes []byte, c Client) {
			log.Debugf("onBinary:%s", string(bytes))
		},
	}, nil
}

func (c Client) Close() {
	c.conn.Close()
}

func (c Client) StartReadMessage() {
	for {
		msgType, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Error("read message error: ", err)
			continue
		}

		switch msgType {
		case websocket.TextMessage:
			c.onText(string(message), c)
		case websocket.BinaryMessage:
			c.onBinary(message, c)
		}
	}
}

func (c *Client) MountBinaryHandler(f OnBinary) {
	if f != nil {
		c.onBinary = f
	}
}
