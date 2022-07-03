package proxy

import (
	"github.com/fzdwx/burst/pkg/logx"
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"net"
)

type (
	UserConn struct {
		// Id uuid
		Id   string
		conn net.Conn
		c    *Container
		// key used to identify the service started by the server
		key string
	}
)

func NewUserConn(conn net.Conn, c *Container, key string) *UserConn {
	return &UserConn{
		Id:   uuid.New().String(),
		conn: conn,
		c:    c,
		key:  key,
	}
}

// ReadUserRequest read user request to client with to write intranet service.
func (u UserConn) ReadUserRequest() {
	for {
		// todo read buffer size
		buf := make([]byte, 1024)
		n, err := u.conn.Read(buf)
		if err != nil {
			u.err(err).Msg("read user connection")
			break
		}

		data, err := u.buildUserRequest(buf, n)

		if err != nil {
			u.err(err).Msg("encode userRequest")
			continue
		}

		// todo write client receive
		u.c.WriteBinary(data)
	}
}

func (u UserConn) buildUserRequest(buf []byte, n int) ([]byte, error) {
	return protocal.NewUserRequest(buf[:n], u.key, u.Id).Encode()
}

func (u UserConn) err(err error) *zerolog.Event {
	return logx.Err(err).Str("connId", u.Id).Str("key", u.key)
}
