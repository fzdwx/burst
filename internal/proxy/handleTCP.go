package proxy

import (
	"github.com/fzdwx/burst/internal/logx"
	"io"
	"net"
	"strings"
)

// handleTCP
// each time a new port is started to handle the corresponding proxy request
func (c *Container) handleTCP(info *internal.ServerProxyInfo) (error, *internal.ClientProxyInfo, io.Closer) {
	tcp, err := net.ListenTCP(info.ChannelType, nil)
	if err != nil {
		return err, nil, nil
	}

	c.AddCloser(tcp)
	serverPort := tcp.Addr().(*net.TCPAddr).Port

	cp := &internal.ClientProxyInfo{
		ChannelType:  info.ChannelType,
		IntranetAddr: info.Addr,
		ServerPort:   serverPort,
	}
	info.ClientProxyInfo = cp
	info.BindListener = tcp

	go func() {
		for {
			// accept user connection
			conn, err := tcp.AcceptTCP()
			if err != nil {
				if strings.ContainsAny("use of closed network connection", err.Error()) || strings.ContainsAny("EOF", err.Error()) {
					return
				}

				logx.Err(err).Str("channelType", info.ChannelType).Msg("accept user connection")
				return
			}

			userConn := NewUserConn(conn, c, cp.Key())
			c.AddCloser(conn)
			c.AddUserConn(userConn)
			clean := c.CleanUserConn(userConn)

			err = userConn.OnUserConnect()
			if err != nil {
				clean()
				continue
			}

			info.BindUserConn = append(info.BindUserConn, userConn.conn)
			go userConn.StartRead(clean)
			go userConn.StartWrite(clean)
		}
	}()

	return nil, cp, tcp
}
