package proxy

import (
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"io"
	"net"
)

type (
	UserConn struct {
		// Id uuid
		Id   string
		conn net.Conn
		c    *Container
		// key used to identify the service started by the server
		key       string
		writeChan chan []byte
	}
)

func NewUserConn(conn net.Conn, c *Container, key string) *UserConn {
	return &UserConn{
		Id:        uuid.New().String(),
		conn:      conn,
		c:         c,
		key:       key,
		writeChan: make(chan []byte),
	}
}

// OnUserConnect notify the client to monitor the intranet service
func (u UserConn) OnUserConnect() error {
	bytes, err := protocal.NewUserConnect(u.key, u.Id).Encode()
	if err != nil {
		u.err(err).Msg("encode userConnect")
		return err
	}

	u.c.WriteBinary(bytes)
	return nil
}

// StartRead read user request to client with to write intranet service.
func (u UserConn) StartRead(clean func()) {
	defer clean()

	for {
		// todo read buffer size
		buf := make([]byte, 1024)
		n, err := u.conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				u.err(err).Msg("read user connection")
			}
			return
		}

		if n == 0 {
			continue
		}

		bytes, err := protocal.NewUserRequest(buf[:n], u.key, u.Id).Encode()

		if err != nil {
			u.err(err).Msg("encode userRequest")
			continue
		}

		u.c.WriteBinary(bytes)
		logx.Debug().Int("write", n).Msg("write to client on user request")
	}
}

// StartWrite start Write intranet response to user
func (u UserConn) StartWrite(clean func()) {
	defer clean()

	for {
		select {
		case data := <-u.writeChan:
			// write user request data to internet
			_, err := u.conn.Write(data)
			if err != nil {
				u.err(err).Msg("write to user")
				return
			}
		}
	}
}

// Write data to user
func (u UserConn) Write(data []byte) {
	u.writeChan <- data
}

func (u UserConn) err(err error) *zerolog.Event {
	return logx.Err(err).Str("connId", u.Id).Str("key", u.key)
}
