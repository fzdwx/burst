package burst

import (
	"github.com/fzdwx/burst/burst-client/common"
	"github.com/fzdwx/burst/burst-client/protocol"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
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
		// proxyInfo mapping: key serverPort ; value ip:port.
		proxyInfo map[int32]*protocol.Proxy
	}

	// OnText is a callback method that will be called back when there is a text type message.
	OnText func(string, *Client)

	// OnBinary is a callback method that will be called back when there is a binary type message.
	OnBinary func([]byte, *Client)
)

// Connect to Server,will return new Client.
func Connect(url url.URL) (*Client, *http.Response, error) {
	log.Printf("start connecting to %s", url.String())
	c, resp, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		return nil, resp, err
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
	}, nil, nil
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

// SetProxyInfo set ports mapping.
func (c *Client) SetProxyInfo(proxyInfo map[int32]*protocol.Proxy) {
	c.proxyInfo = proxyInfo
}

// ProxyInfo get ports mapping.
func (c *Client) ProxyInfo() map[int32]*protocol.Proxy {
	return c.proxyInfo
}

// GetProxy Get the local port(ip:port) corresponding to the server port.
func (c Client) GetProxy(serverExportPort int32) (*protocol.Proxy, bool) {
	v, ok := c.proxyInfo[serverExportPort]
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
