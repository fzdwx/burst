package client

import (
	"fmt"
	"github.com/fzdwx/burst/internal/linereader"
	"github.com/fzdwx/burst/internal/logx"
	"github.com/fzdwx/burst/internal/protocal"
	"github.com/fzdwx/burst/internal/wsx"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"io"
	"net"
	"net/url"
	"os"
	"strings"
	"time"
)

type (
	Client struct {
		*wsx.Wsx
		token      string
		config     Config
		serverAddr string
		serverHost string
		proxy      map[string]internal.ClientProxyInfo
		internet   map[string]*InternetService
	}
)

func NewClient(token string, config Config) *Client {
	return &Client{
		token:      token,
		config:     config,
		serverAddr: internal.FormatAddr(config.Server.Host, config.Server.Port),
		serverHost: config.Server.Host,
		proxy:      make(map[string]internal.ClientProxyInfo),
		internet:   make(map[string]*InternetService),
	}
}

func (c *Client) Token() string {
	return c.token
}

func (c *Client) ServerAddr() string {
	return c.serverAddr
}

// Connect  to server,and init connection
func (c *Client) Connect(init func(wsx *wsx.Wsx)) {
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
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				logx.Fatal().Err(err).Msg("connect to server read resp body")
			}
			logx.Fatal().Msgf("connect to server fail: %s", string(body))
		}
		return
	}

	c.Wsx = wsx.NewClassicWsx(conn)

	if init != nil {
		init(c.Wsx)
	}

	go c.Wsx.StartReading(0)
	go c.Wsx.StartWriteHandler(time.Second * 5)
}

// AddProxy add proxy info to client
//
// do not consider repetition
func (c *Client) AddProxy(a protocal.AddProxy) {
	for _, info := range a.Proxy {
		c.proxy[info.Key()] = info
		logx.Info().Msgf("add proxy: intranet [%s] to [%s]", info.IntranetAddr, info.Address(c.serverHost))
	}
}

// RemoveProxy remove proxy info from client
func (c *Client) RemoveProxy(r protocal.RemoveProxy) {
	for _, info := range r.Proxy {
		delete(c.proxy, info.Key())
		logx.Info().Msgf("remove proxy: intranet [%s] to [%s]", info.IntranetAddr, info.Address(c.serverHost))
	}
}

func (c *Client) GetProxy(key string) (internal.ClientProxyInfo, bool) {
	info, ok := c.proxy[key]
	return info, ok
}

func (c *Client) AddInterNetService(net *InternetService) {
	c.internet[net.connId] = net
}

func (c *Client) GetInternetService(connId string) (*InternetService, bool) {
	internet, ok := c.internet[connId]
	return internet, ok
}

func (c *Client) RemoveInterNetService(interNet *InternetService) {
	delete(c.internet, interNet.connId)
}

func (c *Client) ReaderCommand(f func(line string, client *Client)) {
	lr := linereader.New(os.Stdin)

	fmt.Println("please input command: (u for usage)")
	// Get all the lines
	for {
		fmt.Print("> ")
		select {
		case line := <-lr.Ch:
			f(line, c)
		}
	}
}

type (
	InternetService struct {
		conn      net.Conn
		connId    string
		proxy     internal.ClientProxyInfo
		writeChan chan []byte
	}
)

func NewInternetService(conn net.Conn, connId string, proxy internal.ClientProxyInfo) *InternetService {
	return &InternetService{
		conn:      conn,
		connId:    connId,
		proxy:     proxy,
		writeChan: make(chan []byte),
	}
}

func (s *InternetService) StartRead(c *Client, clean func()) {
	defer clean()

	for {
		// todo read buffer size
		buf := make([]byte, 1024)
		n, err := s.conn.Read(buf)
		if err != nil {
			if strings.ContainsAny("use of closed network connection", err.Error()) || strings.ContainsAny("EOF", err.Error()) {
				return
			}

			s.err(err).Msg("read from internet")
			return
		}

		if n == 0 {
			continue
		}

		bytes, err := protocal.NewIntranetResponse(buf[:n], c.token, s.connId).Encode()
		if err != nil {
			s.err(err).Msg("encode internet response")
			continue
		}

		// write to server with forward to user
		c.WriteBinary(bytes)
	}
}

func (s *InternetService) StartWrite(clean func()) {
	defer clean()

	for {
		select {
		case data := <-s.writeChan:
			// write user request data to internet
			n, err := s.conn.Write(data)
			if err != nil {
				s.err(err).Msg("write to internet")
				return
			}
			logx.Debug().Int("write", n).Msg("write user request data to internet")
		}
	}
}

// CloseInternet close internet service
func (c *Client) CloseInternet(proxy []internal.ClientProxyInfo) {
	for _, service := range c.internet {
		for _, p := range proxy {
			if service.proxy.Key() == p.Key() {
				service.conn.Close()
				logx.Debug().Msgf("close internet connection: %s", p.Key())
			}
		}
	}
}

func (s *InternetService) err(err error) *zerolog.Event {
	return logx.Err(err).Str("internet", s.proxy.IntranetAddr).Str("connId", s.connId)
}

func (s *InternetService) Write(data []byte) {
	s.writeChan <- data
}
