package cs

import (
	"github.com/fzdwx/burst/burst-client/protocol"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
)

type (
	Client struct {
		conn     *websocket.Conn
		token    string
		onText   OnText
		onBinary OnBinary
		// key serverPort v localPort
		ports map[int32]int32
	}

	OnText func(string, *Client)

	OnBinary func([]byte, *Client)
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
		onText: func(s string, c *Client) {
			log.Debugf("onText:%s", s)
		},
		onBinary: func(bytes []byte, c *Client) {
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
			c.onText(string(message), &c)
		case websocket.BinaryMessage:
			c.onBinary(message, &c)
		}
	}
}

func (c *Client) MountBinaryHandler(f OnBinary) {
	if f != nil {
		c.onBinary = f
	}
}

func (c *Client) SetPorts(ports map[int32]int32) {
	c.ports = ports
}

func (c *Client) Ports() map[int32]int32 {
	return c.ports
}

func (c Client) Over(err error) {
	log.Error("user connect close cause: ", err)
	c.Close()
	os.Exit(1)
}

func (c Client) LocalPort(serverExportPort int32) (int32, bool) {
	v, ok := c.ports[serverExportPort]
	return v, ok
}

func (c Client) Write(userConnectId string, bytes []byte) error {
	return c.conn.WriteMessage(websocket.BinaryMessage, protocol.Encode(userConnectId, bytes))
}
