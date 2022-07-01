package client

import (
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/url"
)

type Client struct {
	conn       websocket.Conn
	token      string
	config     Config
	serverAddr string
}

func NewClient(token string, config Config) *Client {
	return &Client{
		token:      token,
		config:     config,
		serverAddr: burst.FormatAddr(config.Server.Host, config.Server.Port),
	}
}

func (c *Client) Connect() {
	u := url.URL{
		Scheme:   "ws",
		Host:     c.serverAddr,
		Path:     "accept",
		RawQuery: "token=" + c.token,
	}

	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		if resp.Body != nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logx.Fatal().Err(err).Msg("connect to server read resp body")
			}
			logx.Fatal().Msgf("connect to server fail: %s", string(body))
		}
		return
	}

	// test
	conn.WriteMessage(websocket.TextMessage, []byte("123"))
}