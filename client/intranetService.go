package client

import (
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/rs/zerolog"
	"net"
)

type (
	InternetService struct {
		conn      net.Conn
		connId    string
		proxy     pkg.ClientProxyInfo
		writeChan chan []byte
	}
)

func NewInternetService(conn net.Conn, connId string, proxy pkg.ClientProxyInfo) *InternetService {
	return &InternetService{
		conn:      conn,
		connId:    connId,
		proxy:     proxy,
		writeChan: make(chan []byte),
	}
}

func (s InternetService) StartRead(c *Client) {
	// todo clean
	for {
		// todo read buffer size
		buf := make([]byte, 1024)
		n, err := s.conn.Read(buf)
		if err != nil {
			s.err(err).Msg("read from internet")
			return
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

func (s InternetService) StartWrite() {
	// todo clean
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

func (s InternetService) err(err error) *zerolog.Event {
	return logx.Err(err).Str("internet", s.proxy.IntranetAddr).Str("connId", s.connId)
}

func (s InternetService) Write(data []byte) {
	s.writeChan <- data
}
