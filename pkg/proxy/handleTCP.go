package proxy

import (
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/logx"
	"io"
	"net"
)

func (c *Container) handleTCP(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	tcp, err := net.ListenTCP(info.ChannelType, nil)
	if err != nil {
		return err, nil, nil
	}

	c.AddCloser(tcp)
	serverPort := tcp.Addr().(*net.TCPAddr).Port

	cp := &pkg.ClientProxyInfo{
		ChannelType:  info.ChannelType,
		IntranetAddr: info.Addr,
		ServerPort:   serverPort,
	}

	go func() {
		for {
			// accept user connection
			conn, err := tcp.AcceptTCP()
			if err != nil {
				logx.Err(err).Str("channelType", info.ChannelType).Msg("accept user connection")
				continue
			}

			userConn := NewUserConn(conn, c, cp.Key())
			c.AddCloser(conn)
			c.AddUserConn(userConn)

			err = userConn.OnUserConnect()
			if err != nil {
				continue
			}

			go userConn.StartRead()
			go userConn.StartWrite()
		}
	}()

	return nil, cp, tcp
}