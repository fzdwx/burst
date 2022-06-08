package burst

import (
	"github.com/fzdwx/burst/burst-client/common"
	"github.com/fzdwx/burst/burst-client/protocol"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
)

type (
	// Client related operations and information.
	Client struct {
		//
		conn     *websocket.Conn
		token    string
		onText   OnText
		onBinary OnBinary
		// ports mapping: key serverPort ; value localPort.
		ports map[int32]int32
	}

	// OnText is a callback method that will be called back when there is a text type message.
	OnText func(string, *Client)

	// OnBinary is a callback method that will be called back when there is a binary type message.
	OnBinary func([]byte, *Client)
)

// Connect to Server,will return new Client.
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

// Close the connection to the server.
func (c Client) Close() error {
	return c.conn.Close()
}

// React to message from the server.
func (c Client) React() {
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
		default:
			if common.IsDebug() {
				log.WithFields(log.Fields{
					"msgType": msgType,
					"message": string(message),
				}).Debug("unSupport message")
			}
		}
	}
}

// MountBinaryHandler Mount processing binary data sent from the server.
func (c *Client) MountBinaryHandler(f OnBinary) {
	if f != nil {
		c.onBinary = f
	}
}

// MountTextHandler Mount processing text data sent from the server.
func (c *Client) MountTextHandler(f OnText) {
	if f != nil {
		c.onText = f
	}
}

// SetPorts set ports mapping.
func (c *Client) SetPorts(ports map[int32]int32) {
	c.ports = ports
}

// Ports get ports mapping.
func (c *Client) Ports() map[int32]int32 {
	return c.ports
}

// LocalPort Get the local port corresponding to the server port.
func (c Client) LocalPort(serverExportPort int32) (int32, bool) {
	v, ok := c.ports[serverExportPort]
	return v, ok
}

// Over is called when there is some exception that needs to close the client.
func (c Client) Over(err error) {
	log.Error("stop client cause: ", err)
	c.Close()
	os.Exit(1)
}

// ToServer forward data to server.
//
// The server then routes to the specified user based on the userConnectId.
func (c Client) ToServer(userConnectId string, data []byte) error {
	return c.conn.WriteMessage(websocket.BinaryMessage, protocol.Encode(userConnectId, data, c.token))
}
