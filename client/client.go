package client

import (
	"github.com/fzdwx/burst"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/wsx"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/url"
	"time"
)

type Client struct {
	*wsx.Wsx
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

	// connect to server and check not error
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

	c.Wsx = wsx.NewClassicWsx(conn)

	go c.Wsx.StartReading(time.Second * 20)
	go c.Wsx.StartWriteHandler(time.Second * 5)
}