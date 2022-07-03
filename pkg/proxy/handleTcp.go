package proxy

import (
	"github.com/fzdwx/burst/pkg"
	"github.com/fzdwx/burst/pkg/logx"
	"net"
)

func (c *Container) handleTCP(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo) {
	tcp, err := net.ListenTCP(info.ChannelType, nil)
	if err != nil {
		return err, nil
	}

	c.Closer = append(c.Closer, tcp)
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
			go userConn.ReadUserRequest()
		}
	}()

	return nil, cp
}
